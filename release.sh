OSARCH="windows/amd64 darwin/amd64 linux/amd64"
rm -rf ./release
gox -osarch="${OSARCH}" -output="release/{{.OS}}_{{.Arch}}/tsg"
cd ./release
for i in $OSARCH
do
  zipname=$(echo ${i} | sed -e "s/\//\_/")
  zip -r ${zipname}.zip ${zipname}
  rm -rf ${zipname}
done
