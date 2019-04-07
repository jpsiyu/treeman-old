<template>
    <div class="record">
        <div class="record-origin" @mouseenter="onMouseEnter" @mouseleave="onMouseLeave">
            <span class="record-detail">{{item.detail}}</span>
            <span class="record-comment">{{item.comment}}</span>
            <transition name="slide-fade">
                <div class="record-btns" v-if="showBtns">
                    <button class="btnDel record-btns__btn" @click="deleteRecord">Delete</button>
                    <button
                        class="btnModify record-btns__btn"
                        @click="showModify=!showModify"
                    >Modify</button>
                </div>
            </transition>
        </div>
        <div class="record-modify" v-if="showModify">
            <div class="record-modify__space" @click="onSpaceClick">+</div>
            <span class="record-modify__title">Detail:</span>
            <textarea class="record-modify__area" v-model="detail"></textarea>
            <span class="record-modify__title">Comment:</span>
            <textarea class="record-modify__area" v-model="comment"></textarea>
            <button class="record-modify__btn" @click="onModifyClick">Modify</button>
        </div>
        <span class="record-timestamp">{{item.timestamp | date}}</span>
    </div>
</template>

<script>
import request from "../request";
import Timer from "../lib/timer";
export default {
    props: ["item"],
    data: function() {
        return {
            timer: new Timer(),
            showBtns: false,
            detail: this.item.detail,
            comment: this.item.comment,
            showModify: false
        };
    },
    filters: {
        date: function(timestamp) {
            const d = new Date(timestamp * 1000);
            const s = d.toLocaleDateString();
            return s;
        }
    },
    methods: {
        deleteRecord() {
            request.delRecord(this.item._id).then(response => {
                this.$emit("getRecord");
            });
        },
        onMouseEnter() {
            this.timer.start(_ => {
                if (this.timer.getTimePass() > 1) {
                    this.showBtns = true;
                    this.timer.stop();
                }
            });
        },
        onMouseLeave() {
            this.timer.stop();
            this.showBtns = false;
        },
        onSpaceClick() {
            this.showModify = !this.showModify;
        },
        onModifyClick() {
            request
                .updateRecord(this.item._id, this.detail, this.comment)
                .then(res => {
                    this.showModify = false;
                    this.$emit("getRecord");
                });
        }
    }
};
</script>

<style scoped>
.record {
    padding: 10px;
    border-radius: 5px;
    margin: 5px;
    position: relative;
    background-color: white;
    box-sizing: border-box;
}

.record-detail {
    display: block;
    white-space: pre-line;
    padding: 30px 0 10px 0;
    line-height: 200%;
    letter-spacing: 1.5px;
    font-weight: bold;
}

.record-comment {
    padding: 10px 0 0px 0;
    white-space: pre;
    display: block;
    color: gray;
    white-space: pre-line;
    line-height: 200%;
    letter-spacing: 1.5px;
}

.record-btns {
    position: absolute;
    display: flex;
    flex-direction: column;
    background-color: white;
    top: 0;
    right: 0;
    padding: 10px;
}

.record-btns__btn {
    outline: none;
    width: 80px;
    height: 30px;
    margin: 5px 0;
    cursor: pointer;
}

.btnDel {
    background-color: red;
    color: white;
}

.btnModify {
    background-color: seagreen;
    color: white;
}

.record-modify {
    display: flex;
    flex-direction: column;
}
.record-modify__space {
    text-align: center;
    font-size: 20px;
    padding: 20px 0;
    cursor: pointer;
    color: lightslategray;
}

.record-modify__area {
    color: gray;
    outline: none;
    border: none;
    width: 100%;
    height: 100px;
    box-sizing: border-box;
    font-size: 16px;
    word-spacing: 1.5px;
    padding: 10px;
    background-color: rgba(240, 240, 240, 0.8);
}

.record-modify__title {
    font-size: 16px;
    padding: 5px;
    font-weight: bold;
    color: lightslategray;
    border-radius: 10px;
}

.record-modify__btn {
    outline: none;
    border: none;
    width: 100%;
    height: 50px;
    background-color: green;
    color: white;
    cursor: pointer;
    font-size: 18px;
    font-weight: bold;
}

.record-timestamp {
    display: block;
    margin-top: 20px;
    color: lightseagreen;
    text-align: end;
}

.slide-fade-enter-active {
    transition: all 0.3s ease;
}
.slide-fade-leave-active {
    transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
}
.slide-fade-enter,
.slide-fade-leave-to {
    transform: translateX(10px);
    opacity: 0;
}
</style>
