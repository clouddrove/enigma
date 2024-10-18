FROM python:3.11-slim

# Set working directory
WORKDIR /docs

# Install MkDocs and any required dependencies
RUN pip install mkdocs

# Copy your MkDocs project files into the container
COPY . .

# Expose the default MkDocs development server port
EXPOSE 8000

# Start MkDocs development server
CMD ["mkdocs", "serve", "--dev-addr=0.0.0.0:8000"]