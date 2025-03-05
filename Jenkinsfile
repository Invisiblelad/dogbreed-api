pipeline {
    agent any

    environment {
        IMAGE_REPO = "invisiblelad/dogbreed"
        HELM_VALUES = "dogbreed/values.yaml"
    }

    stages {
        stage('Checkout') {
            steps {
                script {
                    checkout scm
                }
            }
        }

        stage('Get Commit Hash') {
            steps {
                script {
                    COMMIT_HASH = sh(script: "git rev-parse --short HEAD", returnStdout: true).trim()
                    echo "Commit Hash: ${COMMIT_HASH}"
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    sh "docker build -t ${IMAGE_REPO}:${COMMIT_HASH} ."
                }
            }
        }

        stage('Push Image to Repository') {
            steps {
                script {
                    withCredentials([
                        string(credentialsId: 'dockeruser', variable: 'DOCKER_USER'),
                        string(credentialsId: 'dockerpwd', variable: 'DOCKER_PASS')
                    ]) {
                        sh """
                        echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin
                        docker push ${IMAGE_REPO}:${COMMIT_HASH}
                        """
                    }
                }
            }
        }

        stage('Update Helm values.yaml') {
            steps {
                script {
                    sh """
                    sed -i 's|tagV2: .*|tagV2: tagV1:|g' ${HELM_VALUES}
                    sed -i 's|tagV1: .*|tagV1: ${COMMIT_HASH}|g' ${HELM_VALUES}
                    """
                }
            }
        }

        stage('Commit & Push Changes') {
            steps {
                script {
                    withCredentials([usernamePassword(credentialsId: 'github_creds', usernameVariable: 'GIT_USER', passwordVariable: 'GIT_PASS')]) {
                        sh """
                        git config --global user.email "jenkins@example.com"
                        git config --global user.name "Jenkins"

                        # Ensure we're on the right branch & commit changes first
                        git add ${HELM_VALUES}
                        git commit -m "Update Helm values.yaml with new tag ${COMMIT_HASH}"

                        # Securely push changes using credential helper
                        git remote set-url origin https://${GIT_USER}:${GIT_PASS}@github.com/invisiblelad/dogbreed-api.git
                        git checkout feature
                        git push origin feature
                        """
                    }
                }
            }
        }
    }
}
