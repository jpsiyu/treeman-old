<template>
    <div class="page-add" @click="close">
        <transition name="fade">
            <div class="detail" v-if="show" @click.stop>
                <div class="top">
                    <span class="title">添加人物</span>
                    <div class="close noselect" @click="close">X</div>
                </div>
                <form class="form" @submit.prevent="onSubmit">
                    <pair>
                        <template v-slot:key>
                            <label>名字:</label>
                        </template>
                        <template v-slot:value>
                            <input class="inputname" type="text" v-model="name">
                        </template>
                    </pair>

                    <pair>
                        <template v-slot:key>
                            <label for>性别:</label>
                        </template>
                        <template v-slot:value>
                            <div class="radio-wrap">
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
                            <select class="ageselect" v-model="age">
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
        </transition>
    </div>
</template>

<script>
import { MacroGender } from "../macro";
import request from "../request";
import Vue from "vue";
export default {
    name: "PageAdd",
    data: function() {
        return {
            checkErrors: [],
            gender: MacroGender.Male,
            age: 20,
            name: "",
            show: false
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
            this.$emit("pageSwitch", false);
        },
        onSubmit: function(event) {
            this.checkErrors = [];
            if (!this.name) this.checkErrors.push("请输入名字!");
            if (!this.gender) this.checkErrors.push("请选择性别!");

            if (this.checkErrors.length == 0) {
                request
                    .genPerson(this.name, this.age, this.gender)
                    .then(res => {
                        this.$emit("getAllPerson");
                        this.close();
                    });
            }
        }
    },
    mounted: function() {
        this.show = true;
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
    background-color: rgba(0, 0, 0, 0.4);
    margin: auto;
    overflow-y: scroll;
    display: flex;
}

.detail {
    min-height: 250px;
    width: 400px;
    margin: auto;
    background-color: white;
    flex-direction: column;
}

.top {
    background-color: lightseagreen;
    color: white;
    display: flex;
    justify-content: center;
    align-items: center;
    font-weight: bold;
    height: 50px;
    position: relative;
}

.close {
    position: absolute;
    right: 0;
    margin: 0 10px;
    color: white;
    border-radius: 50%;
    width: 30px;
    height: 30px;
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
}

.pair {
    padding: 10px 0;
    display: flex;
    align-items: flex-end;
    justify-content: space-between;
    width: 80%;
    margin: auto;
}

.inputname {
    padding: 3px 5px;
    outline: none;
    border: none;
    background-color: rgba(0, 0, 0, 0.1);
    height: 20px;
    width: 200px;
}

.ageselect {
    width: 210px;
    height: 28px;
    outline: none;
    padding: 3px 5px;
    border: none;
    background-color: rgba(0, 0, 0, 0.1);
    border-radius: none;
}

.radio-wrap {
    width: 200px;
}

.btn {
    padding: 5px 20px;
    outline: none;
    border: none;
    background-color: lightseagreen;
    color: white;
    cursor: pointer;
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

.form {
    padding-top: 20px;
}

.fade-enter {
    opacity: 0.5;
    transform: scale(0, 0);
}
.fade-enter-active {
    transition: opacity 0.5s ease-in, transform 0.5s ease-in;
}
</style>

