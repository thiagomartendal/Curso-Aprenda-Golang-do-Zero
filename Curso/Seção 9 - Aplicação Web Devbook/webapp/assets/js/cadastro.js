async function criarUsuario(evt) {
    evt.preventDefault()

    let nome = document.getElementById('nome').value
    let email = document.getElementById('email').value
    let nick = document.getElementById('nick').value
    let senha = document.getElementById('senha').value
    let confirmarSenha = document.getElementById('confirmar-senha').value

    if (senha != confirmarSenha) {
        // 'As senhas devem ser iguais.'
        return
    } else {
        try {
            const res = await fetch('http://localhost:3000/usuarios', {
                method: 'POST',
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded",
                },
                body: new URLSearchParams({
                    nome,
                    email,
                    nick,
                    senha
                })
            })

            const json = await res.json()
            console.log(json)
        } catch (err) {
            console.log(err.message)
        }
    }
}