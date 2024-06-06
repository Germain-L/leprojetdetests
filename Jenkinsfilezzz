pipeline {
    agent any

    environment {
        GO_VERSION = "1.16"
    }

    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/germain-l/leprojetdetests.git'
            }
        }

        stage('Setup Go') {
            steps {
                sh 'wget https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz'
                sh 'tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz'
                sh 'export PATH=$PATH:/usr/local/go/bin'
                sh 'go version'
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./...'
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
        }
        success {
            echo 'Build and tests succeeded!'
        }
        failure {
            echo 'Build or tests failed.'
        }
    }
}
