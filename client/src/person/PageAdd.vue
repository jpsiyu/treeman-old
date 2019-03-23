<template>
    <div class="page-add">
        <div class="detail">
            <div class="top">
                <h2 class="title">添加人物</h2>
                <div class="close noselect" @click="onClick">X</div>
            </div>
            <form @submit.prevent="onSubmit">
                <div class="pair">
                    <label for="name">名字:</label>
                    <input type="text">
                </div>
                <div class="pair">
                    <label for>性别:</label>
                    <input type="radio" :value="genderM" v-model="genderPick">
                    <label>男</label>
                    <input type="radio" :value="genderF" v-model="genderPick">
                    <label>女</label>
                </div>
                <div class="pair">
                    <label>年龄:</label>
                    <select v-model="age">
                        <option v-for="n in 100" :key="n">{{n}}</option>
                    </select>
                </div>
                <div class="pair">
                    <button class="btn" type="submit">提交</button>
                </div>
            </form>
        </div>
    </div>
</template>

<script>
import { MacroGender } from "../macro";
export default {
    name: "PageAdd",
    data: function() {
        return {
            genderPick: "",
            age: 0
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
        onClick: function() {
            this.$store.commit("hideAdd");
        },
        onSubmit: function(event) {
            console.log("on submit");
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
    padding: 10px;
}

.btn {
    padding: 5px 20px;
    outline: none;
    border: none;
    background-color: blueviolet;
    color: white;
}
</style>

