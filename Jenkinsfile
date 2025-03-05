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
                    withCredentials([usernamePassword(credentialsId: 'github_creds', usernameVariable: 'GIT_USERNAME', passwordVariable: 'GIT_PASSWORD')]) {
                        sh """
                        git config --global user.email "jenkins@example.com"
                        git config --global user.name "Jenkins"

                        # Fetch latest changes
                        git fetch https://${GIT_USERNAME}:${GIT_PASSWORD}@github.com/Invisiblelad/dogbreed-api.git

                        # Stash uncommitted changes
                        git stash || echo "No changes to stash"

                        # Checkout feature branch
                        git checkout feature

                        # Pull latest changes with rebase
                        git pull https://${GIT_USERNAME}:${GIT_PASSWORD}@github.com/Invisiblelad/dogbreed-api.git feature --rebase || (git rebase --abort && exit 1)

                        # Apply stashed changes
                        git stash pop || echo "No stashed changes to apply"

                        # Add modified values.yaml
                        git add ${HELM_VALUES}

                        # Commit changes
                        git commit -m "Updated helm values.yaml with tag ${COMMIT_HASH} [ci skip]" || echo "No changes to commit"

                        # Push changes
                        git push https://${GIT_USERNAME}:${GIT_PASSWORD}@github.com/Invisiblelad/dogbreed-api.git feature
                        """
                    }
                }
            }
        }
    }
}