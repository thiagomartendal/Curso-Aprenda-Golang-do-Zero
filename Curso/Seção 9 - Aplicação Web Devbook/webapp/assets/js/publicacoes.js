async function criarPublicacao(evt) {
    evt.preventDefault()

    const titulo = document.getElementById('titulo').value
    const conteudo = document.getElementById('conteudo').value

    try {
        const res = await fetch('http://localhost:3000/publicacoes', {
            method: 'POST',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
            },
            body: new URLSearchParams({
                titulo,
                conteudo
            })
        })

        if (res.status == 201)
            window.location = '/home'
    } catch (error) {
        console.log(error.message)
    }
}

async function curtida(evt, id) {
    let link = evt.target
    if (link.classList.contains('curtir-publicacao')) {
        await curtirPublicacao(link, id)
    } else {
        await descurtirPublicacao(link, id)
    }
}

async function curtirPublicacao(link, id) {
    try {
        const res = await fetch(`http://localhost:3000/publicacoes/${id}/curtir`, {
            method: 'POST',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
            }
        })

        if (res.status == 204) {
            let curtidas = document.getElementById(`curtidas-${id}`)
            let contador = parseInt(curtidas.innerHTML) + 1
            curtidas.innerHTML = contador + " curtidas"
            link.classList.remove('curtir-publicacao')
            link.classList.add('descurtir-publicacao')
            link.innerHTML = 'Descurtir'
        }
    } catch (error) {
        console.log(error.message)
    }
}

async function descurtirPublicacao(link, id) {
    link.setAttribute('aria-disabled', 'true')
    try {
        const res = await fetch(`http://localhost:3000/publicacoes/${id}/descurtir`, {
            method: 'POST',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
            }
        })

        if (res.status == 204) {
            let curtidas = document.getElementById(`curtidas-${id}`)
            let contador = parseInt(curtidas.innerHTML) - 1
            curtidas.innerHTML = contador + " curtidas"
            link.classList.remove('descurtir-publicacao')
            link.classList.add('curtir-publicacao')
            link.innerHTML = 'Curtir'
        }
    } catch (error) {
        console.log(error.message)
    }
    link.setAttribute('aria-disabled', 'false')
}

async function atualizarPublicacao(id) {
    const titulo = document.getElementById('titulo').value
    const conteudo = document.getElementById('conteudo').value

    try {
        const res = await fetch(`http://localhost:3000/publicacoes/${id}`, {
            method: 'PUT',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
            },
            body: new URLSearchParams({
                titulo,
                conteudo
            })
        })

        if (res.status == 204)
            window.location.reload()
    } catch (error) {
        console.log(error.message)
    }
}

async function deletarPublicacao(id) {
    try {
        const res = await fetch(`http://localhost:3000/publicacoes/${id}`, {
            method: 'DELETE',
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