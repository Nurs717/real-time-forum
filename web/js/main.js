console.log("JS Loaded")

const url = "http://127.0.0.1:8080"

var inputForm = document.getElementById("inputForm")

inputForm.addEventListener("submit", (e) => {

    //prevent auto submission
    e.preventDefault()

    const formdata = new FormData(inputForm)
    fetch(url, {

        method: "POST",
        body: JSON.stringify({ post: formdata.get("post") }),
    }).then(
        response => response.text()
    ).then(
        (data) => {
            console.log(data);
            document.getElementById("serverPostBox").innerHTML = data
        }
    ).catch(
        error => console.error(error)
    )
})