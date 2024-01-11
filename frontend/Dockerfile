FROM node:21-alpine3.18
WORKDIR /app/client

COPY ./  ./

# Clean install all node modules
RUN npm ci --legacy-peer-deps

# Build SvelteKit app
RUN npm run build

# Delete source code files that were used to build the app that are no longer needed
RUN rm -rf src/ static/ emailTemplates/ docker-compose.yml
CMD ["node","build"]