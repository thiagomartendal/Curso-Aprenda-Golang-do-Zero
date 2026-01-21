async function pararDeSeguir(id) {
    try {
        const res = await fetch(`http://localhost:3000/usuarios/${id}/parar-de-seguir`, {
            method: 'POST',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
            }
        })

        if (res.status == 204)
            window.location.reload()
    } catch (error) {
        console.log(error.message)
    }
}

async function seguir(id) {
    try {
        const res = await fetch(`http://localhost:3000/usuarios/${id}/seguir`, {
            method: 'POST',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
            }
        })

        if (res.status == 204)
            window.location.reload()
    } catch (error) {
        console.log(error.message)
    }
}

async function atualizarDados(evt) {
    evt.preventDefault()

    const nome = document.getElementById('nome').value
    const email = document.getElementById('email').value
    const nick = document.getElementById('nick').value

    try {
        const res = await fetch(`http://localhost:3000/editar-usuario`, {
            method: 'PUT',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
            },
            body: new URLSearchParams({
                nome,
                email,
                nick
            })
        })

        if (res.status == 204)
            window.location.reload()
    } catch (error) {
        console.log(error.message)
    }
}

async function atualizarSenha(evt) {
    evt.preventDefault()

    const senhaAtual = document.getElementById('senha-atual').value
    const novaSenha = document.getElementById('nova-senha').value
    const confirmarSenha = document.getElementById('confirmar-senha').value

    if (novaSenha == confirmarSenha) {
        try {
            const res = await fetch(`http://localhost:3000/atualizar-senha`, {
                method: 'POST',
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded",
                },
                body: new URLSearchParams({
                    atual: senhaAtual,
                    nova: novaSenha
                })
            })

            if (res.status == 204)
                window.location.reload()
        } catch (error) {
            console.log(error.message)
        }
    }
}

async function deletarConta(evt) {
    evt.preventDefault()

    try {
        const res = await fetch(`http://localhost:3000/deletar-usuario`, {
            method: 'DELETE',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
            }
        })

        if (res.status == 204)
            window.location = '/logout'
    } catch (error) {
        console.log(error.message)
    }
}