<template>
    <div class="page-add">
        <div class="detail">
            <div class="top">
                <h2 class="title">添加人物</h2>
                <div class="close noselect" @click="close">X</div>
            </div>
            <form @submit.prevent="onSubmit">
                <pair>
                    <template v-slot:key>
                        <label>名字:</label>
                    </template>
                    <template v-slot:value>
                        <input type="text" v-model="name">
                    </template>
                </pair>

                <pair>
                    <template v-slot:key>
                        <label for>性别:</label>
                    </template>
                    <template v-slot:value>
                        <div>
                            <input type="radio" :value="genderM" v-model="gender">
                            <label>男</label>
                            <input type="radio" :value="genderF" v-model="gender">
                            <label>女</label>
                        </div>
                    </template>
                </pair>

                <pair>
                    <template v-slot:key>
                        <label>年龄:</label>
                    </template>
                    <template v-slot:value>
                        <select v-model="age">
                            <option disabled :value="age">{{age}}</option>
                            <option v-for="n in 100" :key="n">{{n}}</option>
                        </select>
                    </template>
                </pair>

                <pair>
                    <template v-slot:key>
                        <div class="error" v-if="checkErrors.length != 0">{{checkErrors[0]}}</div>
                    </template>
                    <template v-slot:value>
                        <button class="btn" type="submit">提交</button>
                    </template>
                </pair>
            </form>
        </div>
    </div>
</template>

<script>
import { MacroGender } from "../macro";
import axios from "axios";
import qs from "qs";
import Vue from "vue";
export default {
    name: "PageAdd",
    props: ["swichPageAddState"],
    data: function() {
        return {
            checkErrors: [],
            gender: "",
            age: 0,
            name: ""
        };
    },
    computed: {
        genderM: function() {
            return MacroGender.Male;
        },
        genderF: function() {
            return MacroGender.Female;
        }
    },
    methods: {
        close: function() {
            this.swichPageAddState(false)
        },
        onSubmit: function(event) {
            this.checkErrors = [];
            if (!this.name) this.checkErrors.push("请输入名字!");
            if (!this.gender) this.checkErrors.push("请选择性别!");

            if (this.checkErrors.length == 0) {
                const ajaxData = {
                    name: this.name,
                    age: this.age,
                    gender: this.gender
                };
                axios
                    .post("/api/genperson", qs.stringify(ajaxData))
                    .then(response => {
                        this.close();
                    });
            }
        }
    },
    components: {
        pair: {
            template: `<div class="pair">
                <div class="pair-key">
                    <slot name="key"></slot>
                </div>
                <div class="pair-value">
                    <slot name="value"></slot>
                </div>
            </div>`
        }
    }
};
</script>

<style scoped>
.page-add {
    position: fixed;
    left: 0;
    top: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(0, 0, 0, 0.2);
    margin: auto;
    overflow-y: scroll;
    display: flex;
}

.detail {
    min-height: 400px;
    width: 600px;
    margin: auto;
    background-color: white;
    flex-direction: column;
}

.top {
    position: relative;
    display: flex;
    justify-content: center;
    font-weight: bold;
    border-bottom: 2px solid gray;
    color: gray;
}

.close {
    position: absolute;
    top: 0;
    right: 0;
    background-color: lightcoral;
    color: white;
    border-radius: 50%;
    width: 35px;
    height: 35px;
    display: flex;
    justify-content: center;
    align-items: center;
    margin: 5px;
}

.pair {
    padding: 15px 0;
    display: flex;
    align-items: flex-end;
    justify-content: space-between;
    width: 60%;
    margin: auto;
}

.btn {
    padding: 5px 20px;
    outline: none;
    border: none;
    background-color: blueviolet;
    color: white;
}

.btn:hover {
    box-shadow: 2px 2px 1px 1px rgba(0, 0, 0, 0.2);
}

.btn:active {
    box-shadow: none;
}

.error {
    color: red;
    font-size: 16px;
    font-weight: bold;
}
</style>

