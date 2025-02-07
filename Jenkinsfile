pipeline {
  agent any
  stages {
    stage('Log Files') {
      parallel {
        stage('Log Files') {
          steps {
            sh 'ls -la'
          }
        }

        stage('Install golang') {
          steps {
            sh '''rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.6.linux-amd64.tar.gz
  && export PATH=$PATH:/usr/local/go/bin && go version'''
          }
        }

      }
    }

  }
}