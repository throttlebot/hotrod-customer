node {
    def scmVars = checkout scm
    stage('Build') {
        def image = docker.build("willwangkelda/hotrod-customer:${scmVars.GIT_COMMIT}")
        image.push()
    }
    stage('Test') {
        // Skip
    }
    stage('Deploy') {
        // Skip
    }
}