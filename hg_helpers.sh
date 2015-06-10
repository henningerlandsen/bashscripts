hgr() {
	dir=${PWD##*/}
	echo -e "\n\033[1;33m$dir:\033[0m"
	hg $1 $2 $3 $4 $5
	while read p; do
		dir=${p:41}
		cd $dir
		if [ $? -eq 0 ]; then
			echo -e "\n\033[1;33m$dir:\033[0m"
			hg $1 $2 $3 $4 $5
			cd - > /dev/null
		fi
	done <.hgsubstate
}
