file=cmd/apps/helm.go

sed -i 's|\<d\>|h|g' $file
sed -i 's|docker|helm|g' $file
sed -i 's|Docker|Helm|g' $file

cat $file
