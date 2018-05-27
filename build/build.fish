#!/usr/bin/env fish


# remove old files
for file in (ls)
	if [ $file != 'build.fish' ]
		rm $file
	end
end

# set some variables detailing GOOS/GOARCH targets
set oslist linux darwin windows
set archlist 386 amd64
set mpw_path github.com/deanveloper/mpassgo/cli/mpw

# build to all the compile targets
for os in $oslist
	for ar in $archlist
		env GOOS=$os GOARCH=$ar go build -o mpw-"$os"_"$ar" $mpw_path
	end
end

# move to more human-friendly names
for file in (ls mpw*)
	set new_name $file
	set new_name (string replace 'darwin' 'mac' $new_name)
	set new_name (string replace '386' '32' $new_name)
	set new_name (string replace 'amd64' '64' $new_name)
	switch $new_name
	case "*windows*"
		set new_name "$new_name".exe
	end
	
	if [ $file != $new_name ]	
		mv $file $new_name
	end
end
	

