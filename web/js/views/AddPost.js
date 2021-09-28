import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
    constructor() {
        super();
        this.setTitle("AddPost");
    }

    async getHtml() {
        return `
        <h1>Creat Post</h1>
        <form id="inputForm" onSubmit="return false;">
        <label for="post">Post</label>
        <input id="post" name="post" value="hello!!" type="text">
        <button type="submit" style="width:100px;">Go ...</button>
        </form>
        `;
    }

    async Init() {
        const url = "http://localhost:8080/add"

        var inputForm = document.getElementById("inputForm")

        inputForm.addEventListener("submit", (e) => {

            //prevent auto submission
            e.preventDefault()

            const formdata = new FormData(inputForm)
            fetch(url, {

                method: "POST",
                body: JSON.stringify({ post: formdata.get("post") }),
            }).catch(
                error => console.error(error)
            )
        })
    }
}