pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                sh 'make install-proto'
                sh 'make build'
            }
        }
        stage('Test') {
            steps {
                sh 'make tests'
            }
        }
    }
}