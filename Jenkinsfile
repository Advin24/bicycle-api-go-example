@Library('my-awesome-jenkins-lib@master') 
// import com.asynchrony.checkout.Checkout
// import com.asynchrony.shell.Shell
import com.asynchrony.*

pipeline {
    agent { 
        kubernetes {
            label 'not-delivery-engineers-dind'
            defaultContainer 'dind'
        }
    }
    options {
        skipDefaultCheckout(true)
    }
    stages {
        stage('Docker Info') {
            steps{
                script {
                    hello.call "Brandon"
                    hello.respond()
                }
                sh "docker -v" 
                sh "docker info" 
            }
        }
        stage('Checkout Code') {
            steps{
                container('jnlp') {
                    script {
                        log.info "Checking out master"

                        def checkout = new checkout.Checkout()
                        checkout.from("master","https://gitlab.asynchrony.com", "brandon.adams", "bicycle-api-go")

                        log.warning "Checked out master"
                    }
                }
            }
        }
        stage('Build My Docker Image') {
            steps {
                script {
                    def shell = new shell.Shell(this)
                    
                    shell.noOut("ls -la")
                    shell.noOut("cat Dockerfile")

                    sh "docker build -t bicycle-api-go:v0.0.1 ." 
                }
            }
        }
    }
}