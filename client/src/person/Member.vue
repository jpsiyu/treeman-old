<template>
    <router-link :to="{path: '/record', query:{id: info._id}}">
        <div class="member">
            <div class="background">
                <div class="grid" :class="setColor"></div>
                <div class="grid"></div>
            </div>
            <div class="detail">
                <div class="base">
                    <img class="head" :src="info.gender | imagePath">
                    <p class="name">{{info.name}}</p>
                </div>
                <div class="desc">{{info.desc | desc}}</div>
            </div>
        </div>
    </router-link>
</template>

<script>
import { MacroGender } from "../macro";
export default {
    name: "Member",
    props: ["info"],
    computed: {
        setColor: function() {
            const res = {
                blue: this.info.gender == MacroGender.Male,
                pink: this.info.gender == MacroGender.Female
            };
            return res;
        }
    },
    filters: {
        desc: function(originDesc) {
            if (!originDesc) return "天机不可泄漏";
            else return originDesc;
        },
        imagePath: function(gender) {
            return gender == MacroGender.Male
                ? require("../assets/images/man.png")
                : require("../assets/images/woman.png");
        }
    }
};
</script>

<style scoped>
.member {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    background-color: white;
    width: 370px;
    min-width: 370px;
    height: 200px;
    margin: 5px;
    position: relative;
    cursor: pointer;
}

.background {
    display: flex;
    flex-direction: column;
    width: 100%;
    height: 100%;
}

.grid:nth-child(1) {
    flex: 1;
}

.grid:nth-child(2) {
    flex: 2;
}

.detail {
    position: absolute;
    left: 0;
    right: 0;
    bottom: 0;
    top: 0;
    display: flex;
    flex-direction: column;
    padding: 0 10px;
}

.desc {
    flex: 1;
    color: gray;
    padding: 0 10px;
    word-break: break-all;
    overflow-y: hidden;
}

.base {
    flex: 3;
    display: flex;
    align-items: flex-end;
    padding-bottom: 30px;
}

.head {
    width: 100px;
    height: 100px;
    border-radius: 50%;
    border: 3px solid white;
    background-color: white;
}

.name {
    font-weight: bold;
    font-size: 20px;
    padding: 0 20px;
}

.blue {
    background-color: rgb(0, 158, 248);
}

.pink {
    background-color: rgb(233, 67, 136);
}
</style>

