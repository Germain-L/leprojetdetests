pipeline {
    agent any

    environment {
        GO_VERSION = "1.16"
        GOROOT = "/usr/local/go"
        GOPATH = "${env.WORKSPACE}/go"
        PATH = "${env.PATH}:${GOROOT}/bin:${GOPATH}/bin"
    }

    stages {
        stage('Checkout') {
            steps {
                script {
                    git branch: 'main', url: 'https://github.com/Germain-L/leprojetdetests.git'
                }
            }
        }

        stage('Setup Go') {
            steps {
                script {
                    if (!fileExists('/usr/local/go/bin/go')) {
                        sh 'curl -O https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz'
                        sh 'tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz'
                    }
                }
                sh 'go version'
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./... -v -coverprofile=coverage.out'
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o app'
            }
        }
    }

    post {
        always {
            junit '**/TEST-*.xml'
            archiveArtifacts artifacts: 'coverage.out', allowEmptyArchive: true
        }
        success {
            echo 'Build and tests succeeded!'
        }
        failure {
            echo 'Build or tests failed.'
        }
    }
}
