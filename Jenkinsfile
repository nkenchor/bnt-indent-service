pipeline {
    agent any
    tools {
            go 'go.1.19'
    }
    environment {
        GO117MODULE = 'on'
        APP_DIR = 'bnt-indent-service'
		SERVICE_PROFILE = 'dev,elk'
		DO_REINSTALL = 1
		APP_TYPE= 'Service'
    }
    stages {
        stage('Build') {
            steps {
                       echo 'Compiling and building'
                       sh 'go build .'
                  }
        }
        stage('Deploy') {
            steps {
				sh 'sudo bash sytycc.directories ${APP_DIR} ${APP_TYPE}'
				sh 'sudo bash sytycc.ftp ${APP_DIR} ${APP_TYPE}'
				sh 'sudo bash sytycc.systemd ${APP_DIR} ${DO_REINSTALL}'
            }
        }
    }
}