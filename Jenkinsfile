pipeline {
    agent any
    options {
        skipStagesAfterUnstable()
    }
    // Ensure the desired Go version is installed for all stages,
    // using the name defined in the Global Tool Configuration
    tools { 
        go '1.19'
        dockerTool "docker"
    }
    environment{
        registryCredential=credentials("lisportalprod-docker-pwd")
        AZURE_PIPLINE_STATUS    = 'start'
    }
    stages{
        stage('Test') {
            steps {
                println('start testing')
                //script{
                //    sh 'apt-get update'
                //    sh 'apt-get -y install build-essential'
                //    sh "export \n go run main.go &"
                //    sh 'sleep 30s'                   
                //    sh 'go test ./tests -v' 
                //}
            }
        }
        stage('Build') {
            when {
                expression {
                    ("${env.BRANCH_NAME}" == 'main' || "${env.BRANCH_NAME}" == 'staging')
                }
            }
            steps {
                println('start build')
                script{
                    sh 'docker login -u 4693646 -p Yy8m7KJSm@Lm2'
                    if ("${env.BRANCH_NAME}" == 'staging') {
                        if ("${env.JENKINS_URL}".startsWith('http://192.168.10')){
                            sh 'docker build -t 192.168.10.62:6004/vibrant/lis/productCatalog:staging .'
                        } else if ("${env.JENKINS_URL}".startsWith('http://192.168.60')){
                            sh 'docker build -t 192.168.60.10:6004/vibrant/lis/productCatalog:staging .'
                        } else {
                            error "jenkins url match error: ${env.JENKINS_URL}"
                        }
                        try{
                            //sh 'docker login -u listestdocker -p ${registryCredential} http://listestdocker.azurecr.io'
                            sh 'docker login -u lisportalprod -p ${registryCredential} http://lisportalprod.azurecr.io'
                            sh 'docker build -t lisportalprod.azurecr.io/vibrant/lis/productCatalog:staging .'
                            AZURE_PIPLINE_STATUS = 'build'
                        }catch (Exception e) {
                            echo 'Exception occurred: ' + e.toString()
                        }
                    } else {
                        if ("${env.JENKINS_URL}".startsWith('http://192.168.10')){
                            sh 'docker build -t 192.168.10.62:6004/vibrant/lis/productCatalog:latest .'
                        } else if ("${env.JENKINS_URL}".startsWith('http://192.168.60')){
                            sh 'docker build -t 192.168.60.10:6004/vibrant/lis/productCatalog:latest .'
                        } else {
                            error "jenkins url match error: ${env.JENKINS_URL}"
                        }

                        try{
                            //sh 'docker login -u listestdocker -p ${registryCredential} http://listestdocker.azurecr.io'
                            sh 'docker login -u lisportalprod -p ${registryCredential} http://lisportalprod.azurecr.io'
                            sh 'docker build -t lisportalprod.azurecr.io/vibrant/lis/productCatalog:latest .'
                            AZURE_PIPLINE_STATUS = 'build'
                        }catch (Exception e) {
                            echo 'Exception occurred: ' + e.toString()
                        }

                    }
                }
                println('build finish')
            }
        }
        stage ('Push'){
            when {
                expression {
                    ("${env.BRANCH_NAME}" == 'main' || "${env.BRANCH_NAME}" == 'staging')
                }
            }
            steps{
                println('start push')
                script{
                    if ("${env.BRANCH_NAME}" == 'staging') {
                        if ("${env.JENKINS_URL}".startsWith('http://192.168.10')){
                            sh 'docker push 192.168.10.62:6004/vibrant/lis/productCatalog:staging'
                        } else {
                            sh 'docker push 192.168.60.10:6004/vibrant/lis/productCatalog:staging'
                        }
                        //push clound images 
                        //sh 'docker login -u listestdocker -p ${registryCredential} http://listestdocker.azurecr.io'
                        try{
                            sh 'docker push lisportalprod.azurecr.io/vibrant/lis/productCatalog:staging'
                            AZURE_PIPLINE_STATUS = 'push'
                        }catch (Exception e) {
                            echo 'Exception occurred: ' + e.toString()
                        }
                    } else {
                        if ("${env.JENKINS_URL}".startsWith('http://192.168.10')){
                            sh 'docker push 192.168.10.62:6004/vibrant/lis/productCatalog:latest'
                        } else {
                            sh 'docker push 192.168.60.10:6004/vibrant/lis/productCatalog:latest'
                        }
                        //push clound images 
                        //sh 'docker login -u listestdocker -p ${registryCredential} http://listestdocker.azurecr.io'
                        try{
                            sh 'docker push lisportalprod.azurecr.io/vibrant/lis/productCatalog:latest'
                            AZURE_PIPLINE_STATUS = 'push'
                        }catch (Exception e) {
                            echo 'Exception occurred: ' + e.toString()
                        }
                    }
                }
                println('push finish')
            }
        }
        stage('Deploy'){
            when {
                expression {
                    ("${env.BRANCH_NAME}" == 'main' || "${env.BRANCH_NAME}" == 'staging')
                }
            }
            steps{
                println('deploy')
                script{
                    if ("${env.BRANCH_NAME}" == 'staging') {
                        if ("${env.JENKINS_URL}".startsWith('http://192.168.10')){
                            sh 'ssh lis_updater@192.168.10.212 kubectl rollout restart deployment/lis-productCatalog-deployment-staging'
                        } else {
                            sh 'ssh yuxuan@192.168.60.6 kubectl rollout restart deployment/lis-productCatalog-deployment-staging'
                        }

                        //rollout restart aks deployment
                        try{
                            withKubeConfig([credentialsId:'lisportalprod-kube-config',serverUrl:'https://lisportalprod-dns-nrpmbcaa.hcp.westus3.azmk8s.io:443']){
                                sh 'command -v kubectl |curl -LO "https://storage.googleapis.com/kubernetes-release/release/v1.20.5/bin/linux/amd64/kubectl"'
                                sh 'chmod u+x ./kubectl'
                                sh "./kubectl rollout restart deployment/lis-productCatalog-deployment-staging -n productCatalog"
                            }
                            AZURE_PIPLINE_STATUS = 'deploy'
                        }catch (Exception e) {
                            echo 'Exception occurred: ' + e.toString()
                        }    
                    } else {
                        if ("${env.JENKINS_URL}".startsWith('http://192.168.10')){
                            sh 'ssh wang@192.168.10.212 kubectl rollout restart deployment/lis-productCatalog-deployment'
                        } else {
                            sh 'ssh yuxuan@192.168.60.6 kubectl rollout restart deployment/lis-productCatalog-deployment'
                        }

                        //rollout restart aks deployment
                        try{
                            withKubeConfig([credentialsId:'lisportalprod-kube-config',serverUrl:'https://lisportalprod-dns-nrpmbcaa.hcp.westus3.azmk8s.io:443']){
                                sh 'command -v kubectl |curl -LO "https://storage.googleapis.com/kubernetes-release/release/v1.20.5/bin/linux/amd64/kubectl"'
                                sh 'chmod u+x ./kubectl'
                                sh "./kubectl rollout restart deployment/lis-productCatalog-deployment -n productCatalog"
                            }
                            AZURE_PIPLINE_STATUS = 'deploy'
                        }catch (Exception e) {
                            echo 'Exception occurred: ' + e.toString()
                        }    
                    }

                    if ("${env.BRANCH_NAME}" == 'staging') {
                    jiraSendDeploymentInfo environmentId: 'us-staging-1', environmentName: 'us-staging-1', environmentType: 'staging', site: 'vibrantamerica.atlassian.net', serviceIds: ['lis-billing-productCatalog']
                    } else {
                    jiraSendDeploymentInfo environmentId: 'us-prod-1', environmentName: 'us-prod-1', environmentType: 'production', site: 'vibrantamerica.atlassian.net', serviceIds: ['lis-billing-productCatalog']
                    }
                }
                println('deploy finish')
            }
        }
    }
     post {
            success {
                script {
                    if ("${env.BRANCH_NAME}" == 'main' || "${env.BRANCH_NAME}" == 'staging') {
                    notifySlack('SUCCESS')
                    }
                }
            }
            failure {
                script {

                    if ("${env.BRANCH_NAME}" == 'main' || "${env.BRANCH_NAME}" == 'staging') {

                    notifySlack('FAILURE')
                    }
                }
            }
        }
}
def notifySlack(String buildStatus = 'STARTED') {
    // Build status of null means success.
    buildStatus = buildStatus ?: 'SUCCESS'

    def color
    def environmentType
    if ("${env.BRANCH_NAME}" == "main") {
        environmentType = 'production'
    } else if ("${env.BRANCH_NAME}" == "staging") {
        environmentType = 'staging'
    }
    env.GIT_COMMIT_MSG = sh (script: 'git log -1 --pretty=%B ${GIT_COMMIT}', returnStdout: true).trim()
    env.GIT_AUTHOR = sh (script:"git log -1 --pretty=format:'%an'", returnStdout: true).trim()
    if (buildStatus == 'STARTED') {
        color = '#D4DADF'
    } else if (buildStatus == 'SUCCESS') {
        color = '#BDFFC3'
    } else if (buildStatus == 'FAILURE') {
        color = '#FFFE89'
    } else {
        color = '#FF9FA1'
    }

    def azureBuildStatus
    if ("${AZURE_PIPLINE_STATUS}" == "deploy") {
        azureBuildStatus = 'SUCCESS'
    }else{
        azureBuildStatus = 'FAIL'
    }

    def msg = "${buildStatus}: `${env.JOB_NAME}` branch at `${GIT_COMMIT[0..7]}` by ${env.GIT_AUTHOR} on ${environmentType}.\n Azure-Deploy:${azureBuildStatus}. Last commit message: ${env.GIT_COMMIT_MSG} <${env.BUILD_URL}|#${env.BUILD_NUMBER}>"
    //def msg = "${buildStatus}: `${env.JOB_NAME}` branch at `${GIT_COMMIT[0..7]}` by ${env.GIT_AUTHOR} on ${environmentType}. Last commit message: ${env.GIT_COMMIT_MSG} <${env.BUILD_URL}|#${env.BUILD_NUMBER}>"

    slackSend(color: color, message: msg, channel: 'lis-bot')
}