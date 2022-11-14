FROM cosmtrek/air

ENV TZ="Chile/Continental"
ENV LANGUAGE en_US:en  
ENV LC_ALL en_US.UTF-8 
ENV PATH="${PATH}:/home"

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
RUN git clone https://github.com/so-fancy/diff-so-fancy.git /home/
RUN apt-get update && apt-get install -y less
RUN sh -c "$(wget -O- https://github.com/deluan/zsh-in-docker/releases/download/v1.1.2/zsh-in-docker.sh)" -p git \
    -p https://github.com/zsh-users/zsh-autosuggestions \
    -p https://github.com/zsh-users/zsh-completions

WORKDIR /root
RUN echo 'if [[ -r "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh" ]]; then \
    source "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh" \n \
    fi' | cat - .zshrc > temp && mv temp .zshrc && echo '[[ ! -f /tmp/vscode-zsh/.p10k.zsh ]] || source /tmp/vscode-zsh/.p10k.zsh' >> .zshrc   

WORKDIR /usr/src/app

COPY .p10k.zsh /tmp/vscode-zsh/.p10k.zsh
COPY . .
ENTRYPOINT ["air"]