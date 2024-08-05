#!/usr/bin/env bash

function print_help() {
    cat <<-EOF
deploy logging-broker to systemd

Usage:
    ./deploy.sh [role] [zone]

role:
    broker
    api-server

zone:
    alibj
    alibj02
    alibj-pre
    alish
    alish02
    sj02
    sm
    dohko
    dev
    hwgz
EOF
} 

function error(){
    echo -e "\033[31m\033[01m\033[05m[ $1 ]\033[0m"
}

function info(){
    echo -e "\033[32m[ $1 ]\033[0m"
}

function warn(){
    echo -e "\033[1;33m[ $1 ]\033[0m"
}

function alert(){
    echo -e "\033[45;30m $1 \033[0m"
}

function init_systemd() {
    cfg=$1
    kafka=$2

    mkdir -p /usr/local/logging-broker
    \cp release/broker /usr/local/logging-broker/
    \cp -r configs /usr/local/logging-broker/

    kafka_expr="null=null"

    if [ -n "$kafka" ]; then
        kafka_expr='"KAFKA_VERSION=v0.11"'
    fi

    cat > /etc/systemd/system/logging-broker.service <<EOF
[Unit]
StartLimitIntervalSec=300
StartLimitBurst=5
Description=logging-broker
After=network.target cloud-init.service multi-user.target

[Service]
User=root
Restart=always
RestartSec=30
TimeoutStopSec=30
PIDFile=/tmp/logging-broker.pid
Environment=${kafka_expr}
ExecStart=/usr/local/logging-broker/broker --config=/usr/local/logging-broker/configs/${cfg}

[Install]
WantedBy=multi-user.target
EOF
    systemctl daemon-reload
}

function list_infos() {
	alert "常用命令:  "
	echo "----------------------------------"
	alert "systemctl   start  logging-broker"
	alert "systemctl   stop   logging-broker"
	alert "systemctl   status logging-broker"
	alert "journalctl -u logging-broker -f"
}

function start() {
    systemctl enable logging-broker
    systemctl start logging-broker
    systemctl status logging-broker
}

role=$1
case $role in
    "broker")
        bin="cmd/broker/main.go"
        ;;
    "apiserver")
        bin="cmd/apiserver/main.go"
        ;;
    *)
        error "invalid {role} argument"
        print_help
        exit 1
        ;;
esac

region=$2
kafka=""
cfg=""
case $region in
    "sj02")
        cfg=${region}-${role}.yaml
        ;;
    "alibj")
        cfg=${region}-${role}.yaml
        ;;
    "alibj02")
        cfg=${region}-${role}.yaml
        ;;
    "alibj-pre")
        cfg=${region}-${role}.yaml
        ;;
    "alish")
        kafka="v0.11"
        cfg=${region}-${role}.yaml
        ;;
    "alish02")
        kafka="v0.11"
        cfg=${region}-${role}.yaml
        ;;
    "dohko")
        cfg=${region}-${role}.yaml
        ;;
    "dev")
        cfg=${region}-${role}.yaml
        ;;
    "hwgz")
        cfg=${region}-${role}.yaml
        ;;
    "sm")
        kafka="v0.11"
        cfg=${region}-${role}.yaml
        ;;
    *)
        error "invalid {zone} argument"
        print_help
        exit 1
        ;;
esac

info "begin to build"
./scripts/build.sh ${role}

info "begin to init systemd"
init_systemd ${cfg} ${kafka}

list_infos

info "start logging-broker processor by systemd"
start

info "check terminal stdout and stderr"
journalctl -u logging-broker -f