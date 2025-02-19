pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                git branch: 'main', url: 'https://github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be.git'
            }
        }

        stage('Run Tests') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Build Docker Image') {
            steps {
                sh 'docker build -t core-be:latest .'
            }
        }

        stage('Deploy') {
            steps {
                sh '''
                    cd /home/sweetmnstr/ci-cd &&
                    docker-compose down &&
                    docker-compose up -d
                '''
            }
        }
    }
}