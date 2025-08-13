# Docker Instructions for Portfolio App

This guide will help you run your portfolio app using Docker.

## Prerequisites
- Make sure Docker is installed on your system
- Make sure you're in the project directory

## Understanding the Setup

Your app is now configured with:
- **Base path**: `/` (standard root path)
- **Asset paths**: Standard `/assets/...` paths
- **Nginx**: Simple configuration that serves the SPA from root

## Step 1: Build the Docker Image
```bash
docker build -t portfolio-app .
```

This command:
- `docker build` - tells Docker to build an image
- `-t portfolio-app` - gives your image a name (tag)
- `.` - tells Docker to look for the Dockerfile in the current directory

## Step 2: Run the Container
```bash
docker run -p 3000:80 portfolio-app
```

This command:
- `docker run` - tells Docker to run a container
- `-p 3000:80` - maps port 3000 on your computer to port 80 in the container
- `portfolio-app` - the name of the image to run

## Step 3: Access Your App
Open your web browser and go to: `http://localhost:3000`

**Simple and clean!** No more complex routing or base path issues.

## How the Routing Works

1. **Root access** (`/`) → Serves `index.html` directly
2. **All routes** (`/*`) → Served by React Router via `index.html`
3. **Assets** (`/assets/*`) → Served directly from the filesystem
4. **No redirects, no complex path mapping**

## Useful Docker Commands

### Stop the container
```bash
docker ps  # Find the container ID
docker stop <container-id>
```

### Remove the container
```bash
docker rm <container-id>
```

### Remove the image
```bash
docker rmi portfolio-app
```

### View running containers
```bash
docker ps
```

### View all containers (including stopped ones)
```bash
docker ps -a
```

## What This Docker Setup Does

1. **Build Stage**: Uses Node.js to install dependencies and build your React app
2. **Production Stage**: Uses nginx (a lightweight web server) to serve your built files
3. **Simple Routing**: All routes served by React Router via `index.html`
4. **Performance**: Includes caching for static assets
5. **Security**: Adds security headers for production use

## Troubleshooting

### Port Already in Use
If you get a "port already in use" error, try using a different port:
```bash
docker run -p 3001:80 portfolio-app
```

Then access your app at `http://localhost:3001`

### App Not Loading
- Make sure you've rebuilt the Docker image after making changes
- Check that the container is running: `docker ps`
- Verify the build output contains the expected files

### Assets Not Loading (404 errors)
- Make sure you've rebuilt the Docker image after making changes
- Check that the container is running: `docker ps`
- Verify the build output contains the expected asset files

### Common Issues
1. **App not routing**: Usually caused by nginx not serving `index.html` for SPA routes
2. **Assets not loading**: Usually caused by incorrect file paths or missing files
3. **Build errors**: Usually caused by missing dependencies or configuration issues

## Rebuilding After Changes

If you make changes to the nginx configuration or Dockerfile:
1. Stop the running container
2. Rebuild: `docker build -t portfolio-app .`
3. Run again: `docker run -p 3000:80 portfolio-app`

## Why This Approach is Better

- ✅ **Simple**: No complex routing logic
- ✅ **Standard**: Uses conventional web app patterns
- ✅ **Reliable**: Fewer points of failure
- ✅ **Maintainable**: Easy to understand and debug
- ✅ **Portable**: Works the same way in any environment
