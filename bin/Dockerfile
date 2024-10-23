# Use an official Python runtime as a parent image
FROM python:3.9-slim

# Set the working directory in the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Create a Python script
RUN echo 'print("Hello, World!")' > hello.py

# Run the Python script
CMD ["python", "./hello.py"]
