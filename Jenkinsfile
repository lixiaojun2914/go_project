pipeline{
    agent any
    tools{
        go 'go'
    }
    options{
        buildDiscarder(logRotator(numToKeepStr: '10'))
    }
    stages{
        stage("Build"){
            steps{
                sh 'go build'
                echo 'test'
            }
        }
    }
    post{
        success{
            mail to: '982090951@qq.com', subject: 'build', body: 'test'
        }
    }
}