fetch('/api/hello')
    .then(response => response.text())
    .then(data => console.log(data))
    .catch(error => console.error('Error:', error));