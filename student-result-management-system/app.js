import express from 'express';
import path from 'path';

const app = express();
const port = 3000;

app.use(express.static('./public'));
app.use(express.urlencoded({ extended: true }));

app.get('/', (req, res) => {
    res.sendFile(path.resolve(__dirname, './public/index.html'));
});

app.get('/dev', (req, res) => {
    res.sendFile(path.resolve('./public/about.html'));
});


app.post('/login', (req, res) => {
    const { regId, password } = req.body;
    
        res.redirect(`/profile/${regId}`);
    
});

app.get('/profile/:regno', (req, res) => {
// const fetchButton = document.getElementById('fetchButton');
// const studentInfo = document.getElementById('studentInfo');

// fetchButton.addEventListener('click', () => {
const regno = req.params.regno;
console.log(regno);
console.log("a+ ",req.query);
  fetch(`http://127.0.0.1:3000/student/${regno}`)
    .then(response => response.json())
    .then(data => {
      // Store the response data in a variable
      const studentData = data;
      console.log(studentData);
      const user = studentData;
    let marksTableHTML = '<table><tr><th>Subject</th><th>CAT1</th><th>CAT2</th><th>FAT</th></tr>';
    for (const subject in user.marks) {
        marksTableHTML += `
            <tr>
                <td>${user.marks[subject].subject}</td>
                <td>${user.marks[subject].CAT1}</td>
                <td>${user.marks[subject].CAT2}</td>
                <td>${user.marks[subject].FAT}</td>
            </tr>
        `;
    }
    marksTableHTML += '</table>';
    const studentProfileHTML = `
        <!DOCTYPE html>
        <html>
        <head>
            <title>Student Profile</title>
            <style>
                body {
                    font-family: Arial, sans-serif;
                }
                
                table {
                    width: 100%;
                    border-collapse: collapse;
                }
                
                table, th, td {
                    border: 1px solid black;
                }
                
                th, td {
                    padding: 15px;
                    text-align: left;
                }
            </style>
        </head>
        <body>
            <h1>Welcome ${user.studentName}</h1>
            <p>Registration Number: ${user.regno}</p>
            ${marksTableHTML}
        </body>
        </html>
    `;
    res.send(studentProfileHTML);
    })
    .catch(error => {
      console.error('Error:', error);
    });
;
    
});

app.get('/notfound', (req, res) => {
    res.sendFile(path.resolve('./public/notfound.html'));
});

app.listen(port, () => {
    console.log("Server Listening on Port: http://localhost:" + port);
});
