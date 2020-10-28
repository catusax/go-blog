cd web/front
yarn
echo '' > node_modules/antd/es/style/index.less #umi的bug,会无故引入antd样式
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
