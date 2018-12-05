podTemplate(){
    node('not-delivery-engineers-dind') {
        container('dind') {
            stage('Docker Info') {
                sh "docker -v" 
                sh "docker info" 
            }
            stage('Clone Code') {
                git branch: "$BRANCH_NAME", credentialsId: 'github', url: "https://github.com/Advin24/bicycle-api-go-example.git"
            }
            stage('Build My Docker Image') {

                sh "ls -la" 
                sh "cat Dockerfile" 

                sh "docker build -t bicycle-api-go:v0.0.1 ." 
            } 
        }
    }
}