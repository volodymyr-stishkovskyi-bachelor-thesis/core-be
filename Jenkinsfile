pipeline {
    agent any

    environment {
        GO_VERSION = '1.20'
        DOCKER_IMAGE = 'core-be:latest'
        DOCKER_HUB_REPO = 'sweetmnstr/core-be'
        COMPOSE_PATH = '/home/sweetmnstr/ci-cd/thesis'
        COMPOSE_FILE = 'core-be-docker-compose.yml'
    }

    stages {
        stage('Checkout') {
            steps {
                git branch: 'main', url: 'https://github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be.git'
            }
        }

        stage('Setup Go') {
            steps {
                sh 'go version'
            }
        }

        stage('Run Tests') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Build Docker Image') {
            steps {
                sh '''
                docker build -t ${DOCKER_IMAGE} .
                '''
            }
        }

        stage('Login to Docker Hub') {
            steps {
                withCredentials([string(credentialsId: 'docker-hub-token', variable: 'DOCKER_TOKEN')]) {
                    sh 'echo $DOCKER_TOKEN | docker login -u your-dockerhub-username --password-stdin'
                }
            }
        }

        stage('Push Docker Image') {
            steps {
                sh '''
                docker tag ${DOCKER_IMAGE} ${DOCKER_HUB_REPO}:latest
                docker push ${DOCKER_HUB_REPO}:latest
                '''
            }
        }

        stage('Deploy') {
            steps {
                sh '''
                cd ${COMPOSE_PATH} &&
                docker-compose -f ${COMPOSE_FILE} pull &&
                docker-compose -f ${COMPOSE_FILE} down &&
                docker-compose -f ${COMPOSE_FILE} up -d
                '''
            }
        }
    }
}