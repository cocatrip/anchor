file=cmd/apps/docker.go

sed -i 's|\<j\>|d|g' $file
sed -i 's|jenkins|docker|g' $file
sed -i 's|Jenkins|Docker|g' $file

cat $file
