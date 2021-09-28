import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
    constructor() {
        super();
        this.setTitle("SignUp");
    }

    async getHtml() {
        return `
        <h1>Please Register</h1>
        <form id="inputForm2" onSubmit="return false;">
        <div>
        <label for="firstname">First Name</label>
        <input id="firstname" name="firstname" type="text">
        </div>
        <div>
        <label for="lastname">Last Name</label>
        <input id="lastname" name="lastname" type="text">
        </div>
        <div>
        <label for="email">Email</label>
        <input id="email" name="email" type="text">
        </div>
        <div>
        <label for="password">Password</label>
        <input id="password" name="" type="text">
        </div>
        <div>
        <button type="submit" style="width:100px;">Register</button>
        </div>
        </form>
        `;
    }

    async signUp() {
        const url = "http://localhost:8080/signup"

        var inputForm = document.getElementById("inputForm2")

        inputForm.addEventListener("submit", (e) => {

            //prevent auto submission
            e.preventDefault()

            const formdata = new FormData(inputForm)
            fetch(url, {

                method: "POST",
                body: JSON.stringify({ firstname: formdata.get("firstname"), lastname: formdata.get("lastname"), email: formdata.get("email"), password: formdata.get("password") }),
            }).catch(
                error => console.error(error)
            )
        })
    }
}

// <
// div class = "container" >

//     <
//     form class = "well form-horizontal"
// action = " "
// method = "post"
// id = "contact_form" >
//     <
//     fieldset >

//     <!-- Form Name -->
//     <
//     legend >
//     <
//     center >
//     <
//     h2 > < b > Registration Form < /b></h
// 2 >
//     <
//     /center> <
//     /legend><br>

// <!-- Text input-->

// <
// div class = "form-group" >
//     <
//     label class = "col-md-4 control-label" > First Name < /label> <
//     div class = "col-md-4 inputGroupContainer" >
//     <
//     div class = "input-group" >
//     <
//     span class = "input-group-addon" > < i class = "glyphicon glyphicon-user" > < /i></span >
//     <
//     input name = "first_name"
// placeholder = "First Name"
// class = "form-control"
// type = "text" >
//     <
//     /div> <
//     /div> <
//     /div>

// <!-- Text input-->

// <
// div class = "form-group" >
//     <
//     label class = "col-md-4 control-label" > Last Name < /label> <
//     div class = "col-md-4 inputGroupContainer" >
//     <
//     div class = "input-group" >
//     <
//     span class = "input-group-addon" > < i class = "glyphicon glyphicon-user" > < /i></span >
//     <
//     input name = "last_name"
// placeholder = "Last Name"
// class = "form-control"
// type = "text" >
//     <
//     /div> <
//     /div> <
//     /div>

// <
// div class = "form-group" >
//     <
//     label class = "col-md-4 control-label" > Department / Office < /label> <
//     div class = "col-md-4 selectContainer" >
//     <
//     div class = "input-group" >
//     <
//     span class = "input-group-addon" > < i class = "glyphicon glyphicon-list" > < /i></span >
//     <
//     select name = "department"
// class = "form-control selectpicker" >
//     <
//     option value = "" > Select your Department / Office < /option> <
//     option > Department of Engineering < /option> <
//     option > Department of Agriculture < /option> <
//     option > Accounting Office < /option> <
//     option > Tresurer 's Office</option> <
//     option > MPDC < /option> <
//     option > MCTC < /option> <
//     option > MCR < /option> <
//     option > Mayor 's Office</option> <
//     option > Tourism Office < /option> <
//     /select> <
//     /div> <
//     /div> <
//     /div>

// <!-- Text input-->

// <
// div class = "form-group" >
//     <
//     label class = "col-md-4 control-label" > Username < /label> <
//     div class = "col-md-4 inputGroupContainer" >
//     <
//     div class = "input-group" >
//     <
//     span class = "input-group-addon" > < i class = "glyphicon glyphicon-user" > < /i></span >
//     <
//     input name = "user_name"
// placeholder = "Username"
// class = "form-control"
// type = "text" >
//     <
//     /div> <
//     /div> <
//     /div>

// <!-- Text input-->

// <
// div class = "form-group" >
//     <
//     label class = "col-md-4 control-label" > Password < /label> <
//     div class = "col-md-4 inputGroupContainer" >
//     <
//     div class = "input-group" >
//     <
//     span class = "input-group-addon" > < i class = "glyphicon glyphicon-user" > < /i></span >
//     <
//     input name = "user_password"
// placeholder = "Password"
// class = "form-control"
// type = "password" >
//     <
//     /div> <
//     /div> <
//     /div>

// <!-- Text input-->

// <
// div class = "form-group" >
//     <
//     label class = "col-md-4 control-label" > Confirm Password < /label> <
//     div class = "col-md-4 inputGroupContainer" >
//     <
//     div class = "input-group" >
//     <
//     span class = "input-group-addon" > < i class = "glyphicon glyphicon-user" > < /i></span >
//     <
//     input name = "confirm_password"
// placeholder = "Confirm Password"
// class = "form-control"
// type = "password" >
//     <
//     /div> <
//     /div> <
//     /div>

// <!-- Text input-->
// <
// div class = "form-group" >
//     <
//     label class = "col-md-4 control-label" > E - Mail < /label> <
//     div class = "col-md-4 inputGroupContainer" >
//     <
//     div class = "input-group" >
//     <
//     span class = "input-group-addon" > < i class = "glyphicon glyphicon-envelope" > < /i></span >
//     <
//     input name = "email"
// placeholder = "E-Mail Address"
// class = "form-control"
// type = "text" >
//     <
//     /div> <
//     /div> <
//     /div>

// <!-- Text input-->

// <
// div class = "form-group" >
//     <
//     label class = "col-md-4 control-label" > Contact No. < /label> <
//     div class = "col-md-4 inputGroupContainer" >
//     <
//     div class = "input-group" >
//     <
//     span class = "input-group-addon" > < i class = "glyphicon glyphicon-earphone" > < /i></span >
//     <
//     input name = "contact_no"
// placeholder = "(639)"
// class = "form-control"
// type = "text" >
//     <
//     /div> <
//     /div> <
//     /div>

// <!-- Select Basic -->

// <!-- Success message -->
// <
// div class = "alert alert-success"
// role = "alert"
// id = "success_message" > Success < i class = "glyphicon glyphicon-thumbs-up" > < /i> Success!.</div >

//     <!-- Button -->
//     <
//     div class = "form-group" >
//     <
//     label class = "col-md-4 control-label" > < /label> <
//     div class = "col-md-4" > < br >
//     &
//     nbsp & nbsp & nbsp & nbsp & nbsp & nbsp & nbsp & nbsp & nbsp & nbsp & nbsp & nbsp & nbsp & nbsp < button type = "submit"
// class = "btn btn-warning" > & nbsp & nbsp & nbsp & nbsp & nbsp & nbsp & nbsp & nbspSUBMIT < span class = "glyphicon glyphicon-send" > < /span>&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp&nbsp</button >
//     <
//     /div> <
//     /div>

// <
// /fieldset> <
// /form> <
// /div> <
// /div><!-- /.container -->