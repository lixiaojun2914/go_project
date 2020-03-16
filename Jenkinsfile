pipeline{
    agent any
    stages{
        stage("Build"){
            steps{
                echo "Lixiaojun World"
            }
        }
    }
    post{
        success{
            mail to: '982090951@qq.com', subject: 'build', body: 'test'
        }
    }
}