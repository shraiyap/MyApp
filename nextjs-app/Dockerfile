FROM node:20-alpine
WORKDIR /nextjs-app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build
EXPOSE 9090
CMD ["npm", "start"]