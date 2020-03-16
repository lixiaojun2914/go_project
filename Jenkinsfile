pipeline{
    agent any
    tools{
        go 'go1.2'
    }
    options{
        buildDiscarder(logRotator(numToKeepStr: '10'))
    }
    stages{
        stage("Build"){
            steps{
                sh 'go build'
            }
        }
    }
    post{
        success{
            mail to: '982090951@qq.com', subject: 'build', body: 'test'
        }
    }
}