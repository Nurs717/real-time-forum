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
        <label for="post_body">Post</label>
       
        <div class="filter">
            <span> Post category:</span>
            <select  name="category" >
               
            <option value="sport">sport</option>
            <option value="religion">religion</option>
            <option value="politics">politics</option>
            <option value="science">science</option>
            <option value="others">others</option>
            
            </select>
        </div>
        <input id="post" name="post_body" value="hello!!" type="text">
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
                body: JSON.stringify({ post_body: formdata.get("post_body"), category: formdata.get("category") }),
            }).catch(
                error => console.error(error)
            )
        })
    }
}