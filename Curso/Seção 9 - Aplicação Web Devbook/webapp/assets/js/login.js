async function fazerLogin(evt) {
    evt.preventDefault()

    let email = document.getElementById('email').value
    let senha = document.getElementById('senha').value

    try {
        const res = await fetch('http://localhost:3000/login', {
            method: 'POST',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
            },
            body: new URLSearchParams({
                email,
                senha
            })
        })

        if (res.status == 200)
            window.location = '/home'
        else {
            const json = await res.json()
            console.log(json)
        }
    } catch (error) {
        console.log(error.message)
    }
}