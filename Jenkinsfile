podTemplate(){
    node('not-delivery-engineers-dind') {
        container('dind') {
            stage('Docker Info') {
                sh "docker -v" 
                sh "docker info" 
            }
            stage('Build My Docker Image') {
                sh "ls -la" 
                sh "cat Dockerfile" 

                sh "docker build -t bicycle-api-go:v0.0.1 ." 
            } 
        }
    }
}