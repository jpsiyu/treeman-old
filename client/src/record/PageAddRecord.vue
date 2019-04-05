<template>
    <div class="pageAdd" @click="close">
        <div class="pageAdd-container">
            <span class="pageAdd-areaTitle">Detail:</span>
            <textarea class="pageAdd-area" v-model="detail" @click.stop></textarea>
            <span class="pageAdd-areaTitle">Comment:</span>
            <textarea class="pageAdd-area" v-model="comment" @click.stop></textarea>
            <button class="pageAdd-btn" @click.stop="onBtnSubmit">Submit</button>
            <div class="pageAdd-error" v-if="error!=''">{{error}}</div>
        </div>
    </div>
</template>

<script>
import request from "../request";
export default {
    name: "PageAddRecord",
    props: ["id"],
    data: function() {
        return {
            detail: "",
            comment: "",
            error: ""
        };
    },
    methods: {
        onBtnSubmit: function() {
            if (this.detail == "") {
                this.error = "detail cannot be empty";
                return;
            }
            this.error = "";
            request
                .addRecord(this.id, this.detail, this.comment)
                .then(response => {
                    this.$emit("getRecord");
                    this.close();
                });
        },
        close: function() {
            this.$emit("pageSwitch", false);
        }
    }
};
</script>

<style scoped>
.pageAdd {
    cursor: pointer;
    padding: 20px 0 50px 0;
}

.pageAdd-container {
    background-color: white;
    border-radius: 10px;
    display: flex;
    flex-direction: column;
    align-items: center;
    margin: auto;
    max-width: 500px;
    padding: 5px 0;
}

.pageAdd-area {
    height: 100px;
    width: 100%;
    outline: none;
    box-sizing: border-box;
    padding: 10px;
    font-size: 16px;
    border: none;
    position: relative;
    word-spacing: 1.5px;
    line-height: 150%;
}

.pageAdd-areaTitle {
    font-weight: bold;
    font-size: 16px;
    align-self: flex-start;
    background-color: lightslategrey;
    border-radius: 10px;
    width: 100%;
    color: white;
    padding: 2px 5px;
    box-sizing: border-box;
}

.pageAdd-btn {
    width: 100%;
    height: 40px;
    font-size: 20px;
    background-color: seagreen;
    color: white;
    cursor: pointer;
    outline: none;
}

.pageAdd-error {
    padding: 10px 0;
    color: red;
}
</style>

