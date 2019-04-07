<template>
    <div class="login">
        <transition name="fade">
            <div class="form" v-if="show" @keyup.enter="submit">
                <div class="field">
                    <span>Username</span>
                    <input type="text" v-model="username">
                </div>
                <div class="field">
                    <span>Password</span>
                    <input type="password" v-model="password">
                </div>
                <div class="field">
                    <button @click="submit">Submit</button>
                </div>
                <transition name="err">
                    <span class="error" v-if="error">{{error}}</span>
                </transition>
            </div>
        </transition>
    </div>
</template>

<script>
import request from "../request";
import { MacroServerCode } from "../macro";
export default {
    name: "Login",
    data: function() {
        return {
            show: false,
            error: "",
            timer: null,
            username: "",
            password: ""
        };
    },
    methods: {
        submit: function() {
            if (!this.username) this.setError("Input username");
            else if (!this.password) this.setError("Input password");
            else this.getToken();
        },
        getToken() {
            request.getToken(this.username, this.password).then(serverData => {
                if (serverData.code == MacroServerCode.OK) {
                    this.$store.commit("setTokenStr", serverData.data);
                    this.$router.push("/");
                }else if(serverData.code == MacroServerCode.BusinessErr){
                    this.setError(serverData.data)
                }
            });
        },
        setError: function(msg) {
            this.error = "";
            clearTimeout(this.timer);
            this.error = msg;
            this.timer = setTimeout(_ => {
                this.error = "";
            }, 1000);
        }
    },
    mounted: function() {
        this.show = true;
    }
};
</script>

<style scoped>
.login {
    position: absolute;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
}

.form {
    margin: 100px auto 0;
    width: 400px;
    height: 200px;
    background-color: white;
    display: flex;
    flex-direction: column;
    justify-content: center;
    padding: 30px 50px;
    position: relative;
}

.field {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: space-between;
    color: gray;
    padding: 0 5px;
}

.field span {
    font-weight: bold;
}

.field input {
    outline: none;
    border: none;
    background-color: rgba(229, 236, 240, 1);
    height: 40px;
    width: 250px;
    font-size: 18px;
    padding: 0 10px;
    color: gray;
}

.field button {
    width: 100%;
    height: 40px;
    background-color: lightseagreen;
    color: white;
    cursor: pointer;
    font-size: 18px;
    font-weight: bold;
    outline: none;
    border: none;
    user-select: none;
}

.error {
    background-color: lightcoral;
    color: white;
    padding: 3px 10px;
    border-radius: 10px;
    word-break: break-all;
    position: absolute;
    bottom: 10px;
}

/* transitions */
.fade-enter {
    opacity: 0;
    transform: translateY(200px);
}

.fade-enter-active {
    transition: opacity 1s, transform 1s ease-in;
}

.err-enter,
.err-leave-to {
    opacity: 0;
}
.err-enter-active,
.err-leave-active {
    transition: opacity 0.5s;
}
</style>

