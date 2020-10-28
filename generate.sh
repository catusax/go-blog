cd web/front
yarn
yarn build
cd ../admin
yarn
yarn build
cd ../../
mkdir ./www
cp -r ./web/front/dist/* ./www
cp -r ./web/admin/dist/* ./www
cp -r ./static ./www
cp ./static/*.* ./www

cat > Caddyfile <<EOF
http://localhost/api/* {
	file_server
	reverse_proxy /api/* 127.0.0.1:8080
}

http://localhost {
	file_server
	root $PWD/www
	try_files {path} /index.html
}

http://localhost/admin* {
	file_server
	root $PWD/www
	try_files {path} /index.html
}

http://localhost/static/* {
	file_server
	root $PWD/www
}
EOF