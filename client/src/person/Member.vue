<template>
    <div
        class="member"
        @mouseover="onMouseOver"
        @mouseleave="onMouseLeave"
        @click="onClick"
        @mouseenter="onMouseEnter"
    >
        <div class="member-float" :class="{'member-float--over': over}">
            <img class="member-float__head" :src="info.gender | imagePath">
            <span class="member-float__name">{{info.name}}</span>
            <button class="member-float__del" v-if="showDel" @click.stop="del">Delete</button>
        </div>
    </div>
</template>

<script>
import { MacroGender } from "../macro";
import request from "../request";
import Timer from "../timer";
export default {
    name: "Member",
    props: ["info"],
    data: function() {
        return {
            timer: new Timer(),
            over: false,
            showDel: false,
            clickCount: 0
        };
    },
    methods: {
        del: function() {
            request.delPerson(this.info.name).then(res => {
                this.$emit("getAllPerson");
            });
        },
        onMouseEnter: function() {
            this.timer.start(_ => {
                if (this.timer.getTimePass() > 1) {
                    this.timer.stop();
                    this.showDel = true;
                }
            });
        },
        onMouseLeave: function() {
            this.showDel = false;
        },
        onMouseOver: function() {
            this.over = true;
        },
        onMouseLeave: function() {
            this.over = false;
            this.timer.stop();
            this.showDel = false;
        },
        onClick: function() {
            this.$router.push(`/record?id=${this.info._id}`);
        }
    },
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
    user-select: none;
    width: 100%;
    margin-bottom: 10px;
    border-radius: 10px;
    cursor: pointer;
    position: relative;
    height: 50px;
    margin-bottom: 10px;
}
.member-float {
    background-color: white;
    display: flex;
    align-items: center;
    width: 100%;
    height: 100%;
    border-radius: 10px;
}

.member-float--over {
    position: absolute;
    left: -10px;
    top: -3px;
    box-shadow: 2px 2px 1px rgba(150, 150, 150, 0.5);
}

.member-float__del {
    outline: none;
    color: white;
    background-color: red;
    height: 60%;
    padding: 0 10px;
    position: absolute;
    right: 10px;
    cursor: pointer;
}

.member-float__head {
    height: 80%;
}

.member-float__name {
    color: gray;
    padding: 0px 20px;
}
</style>

