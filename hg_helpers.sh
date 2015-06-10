export hgr_hgexec=hg

testing() {
	export hgr_hgexec=echo
}

__hgr_print_header() {
	echo -e "\n\033[1;33m$1:\033[0m"
}

__hgr_execute_command() {
	eval $1
}

__hgr_recurse_in_subrepos() {
	FILE=.hgsubstate
	if [ -f $FILE ];
	then
		while read p; do
			dir=${p:41}
			cd $dir
			if [ $? -eq 0 ]; then
				__hgr_exec_with_title_command_in_dir $1/$dir $2 ${PWD}
				cd $3
			fi
		done <$FILE
	fi
}

__hgr_exec_with_title_command_in_dir() {
		__hgr_print_header $1
		__hgr_execute_command $2
		__hgr_recurse_in_subrepos $1 $2 $3
}

hgr() {
	dir=${PWD##*/}
	hgcommand="$hgr_hgexec $1 $2 $3 $4 $5"
	__hgr_exec_with_title_command_in_dir $dir $hgcommand ${PWD}
}
