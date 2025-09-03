pipeline {
    agent {
        kubernetes {
            label 'jenkins-agent-golang'
            yaml """
apiVersion: v1
kind: Pod
spec:
  containers:
  - name: golang
    image: golang:1.25.0
    command: ['cat']
    tty: true
    resources:
      requests:
        memory: 1Gi
      limits:
        memory: 2Gi
    env:
    - name: GOPROXY
      value: "https://goproxy.cn,direct"
  - name: jnlp
    image: jenkins/inbound-agent:alpine
    resources:
      requests:
        cpu: 50m
        memory: 128Mi
"""
        }
    }
    parameters {
        booleanParam(name: 'TEST', defaultValue: false, description: '测试标记')
        string(name: 'VERSION', defaultValue: '0.2.1', description: '项目版本号')
        string(name: 'BRANCH', defaultValue: 'r', description: '分支标识')
        string(name: 'GIT_URL', defaultValue: 'https://git.ie8.pub:8443', description: 'git地址')
        choice(name: 'GO_PROXY', choices: ['https://mirrors.aliyun.com/goproxy/,direct', 'https://goproxy.cn,direct'], description: 'Go模块代理')
    }
    environment {
        PROGRAM = "lrc4StreamMusic"
        OS_LIST = "linux darwin windows"
        ARCH_LIST = "arm64 amd64"
        WIN_SUFFIX = ".exe"
        VERSIONS = "v${params.VERSION}.${params.BRANCH}"
        GITEA_PUBLISH_TOKEN = credentials('gitea-publish')
    }
    stages {
        stage('拉取代码') {
            steps {
                container('golang') {
                    checkout scm
                }
            }
        }
        stage('初始化') {
            steps {
                container('golang') {
                    script {
                        // 2. 在步骤内计算动态环境变量
                        // 添加安全目录配置，避免 dubious ownership 错误
                        sh "git config --global --add safe.directory ${env.WORKSPACE}"

                        env.COMMIT = sh(script: 'git rev-parse HEAD', returnStdout: true).trim()
                        env.TIME = sh(script: 'date -u \'+%Y年%m月%d日%H时%M分%S秒\'', returnStdout: true).trim()
                        env.RELEASE_PATH = ${params.GIT_URL}/api/v1/repos/${env.JOB_NAME}/releases
                        echo "开始执行流水线"
                        echo "项目: ${env.PROGRAM}"
                        echo "版本: ${env.VERSIONS}"
                        echo "提交: ${env.COMMIT}"
                        echo "构建时间: ${env.TIME}"
                    }
                }
            }
        }
        stage('设置Go环境') {
            steps {
                container('golang') {
                    script {
                        sh "go env -w GOPROXY=${params.GO_PROXY}"
                        sh "go env | grep GOPROXY"
                        sh "go mod tidy"
                    }
                }
            }
        }
        stage('交叉编译') {
            steps {
                container('golang') {
                    script {
                        def osList = env.OS_LIST.split()
                        def archList = env.ARCH_LIST.split()

                        osList.each { GOOS ->
                            archList.each { GOARCH ->
                                def filename = "${env.PROGRAM}-${GOOS}-${GOARCH}-${env.VERSIONS}"
                                if (GOOS == "windows") {
                                    filename += env.WIN_SUFFIX
                                }

                                echo "${filename}，开始编译"

                                withEnv(["GOOS=${GOOS}", "GOARCH=${GOARCH}" , "CGO_ENABLED=0"]) {
                                    sh """
                                        go build -ldflags " \
                                        -X main.Version=v${params.VERSION}-${params.BRANCH} \
                                        -X main.Commit=${env.COMMIT} \
                                        -X main.BuildTime=${env.TIME} \
                                        -X main.GOOS=${GOOS} \
                                        -X main.GOARCH=${GOARCH}" \
                                        ${params.TEST ? ' -v -x ' : ' '} \
                                        -o ${filename} ./
                                    """
                                }

                                echo "${filename}，编译结束"
                            }
                        }
                    }
                }
            }
        }
        stage('归档制品') {
            steps {
                container('golang') { // 归档操作也在容器内
                    script {
                        sh "mkdir -p artifacts"
                        sh "mv ${env.PROGRAM}-* artifacts/ || true" // 添加 || true 防止找不到文件时失败
                        archiveArtifacts artifacts: 'artifacts/**', fingerprint: true
                    }
                }
            }
        }
        stage('创建 Gitea Release 并上传制品') {
            steps {
                container('golang') {
                    script {
                        // 尝试检查该Tag是否已存在
                        def createReleaseResponse = sh(script: """
                            curl -f -s -H "Authorization: token $GITEA_PUBLISH_TOKEN" \
                            ${env.RELEASE_PATH}/tags/v${params.VERSION}-${params.BRANCH} || true
                        """, returnStdout: true).trim()

                        if (createReleaseResponse != "") {
                            echo "Tag v${params.VERSION}-${params.BRANCH} already exists. Skipping release creation."
                            // 或者在这里添加删除现有Release的逻辑：
                            // sh "curl -X DELETE -H 'Authorization: token \$GITEA_PUBLISH_TOKEN' 'https://git.ie8.pub:8443/api/v1/repos/jiangwe/zulipBot/releases/tags/v${params.VERSION}-${params.BRANCH}'"
                        } else {
                            // 创建 Release (draft: false 表示直接发布)
                            createReleaseResponse = sh(script: """
                                curl -f -s -X POST \
                                    -H "Authorization: token $GITEA_PUBLISH_TOKEN" \
                                    -H "Content-Type: application/json" \
                                    -H "accept: application/json" \
                                    "${env.RELEASE_PATH}" \
                                    -d '{
                                        "tag_name": "v${params.VERSION}-${params.BRANCH}",
                                        "name": "Version ${params.VERSION} (${params.BRANCH})",
                                        "body": "由Jenkins自动构建发布.\n**编译时间:** ${env.TIME}\n**提交代号:** ${env.COMMIT}",
                                        "draft": false,
                                        "prerelease": ${params.BRANCH == 'b' ? 'true' : 'false'}
                                    }'
                            """, returnStdout: true).trim()
                            echo "请求新增发布版本结果消息 ${createReleaseResponse}"
                        }
                        // 2. 解析响应获取 upload_url (Gitea API 返回中包含 id等信息，上传附件是另一个接口)
                        // Gitea API 上传附件的URL格式固定为：https://git.ie8.pub:8443/api/v1/repos/jiangwe/zulipBot/releases/{release_id}/assets?name={filename}
                        // 我们需要从创建返回的JSON中获取release的id
                        def releaseInfo = readJSON text: createReleaseResponse
                        def releaseId = releaseInfo.id
                        def uploadUrl = "${env.RELEASE_PATH}/${releaseId}/assets"

                        // 3. 上传 artifacts 目录下的所有文件作为该 Release 的附件
                        def files = findFiles(glob: 'artifacts/*')
                        for (file in files) {
                            sh """
                                curl -s -X POST \
                                -H "Authorization: token $GITEA_PUBLISH_TOKEN" \
                                -H "accept: application/json" \
                                -H "Content-Type: multipart/form-data" \
                                -F "attachment=@${file.path}" \
                                "${uploadUrl}?name=${file.name}"
                            """
                        }
                    }
                }
            }
        }
    }
    post {
        always {
            echo "流水线执行完成"
            cleanWs()
        }
        success {
            echo "构建成功！"
            // 可添加成功通知，如邮件、钉钉等
        }
        failure {
            echo "构建失败！"
            // 可添加失败告警
        }
    }
}