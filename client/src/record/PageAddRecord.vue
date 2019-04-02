<template>
    <div class="pageAdd" @click="close">
        <div class="pageAdd-container">
            <textarea
                class="pageAdd-area"
                v-model="detail"
                cols="30"
                rows="10"
                @click.stop
                placeholder="detail"
            ></textarea>
            <textarea
                class="pageAdd-area"
                v-model="comment"
                cols="30"
                rows="10"
                @click.stop
                placeholder="comment"
            ></textarea>
            <div class="pageAdd-error" v-if="error!=''">{{error}}</div>
            <button class="pageAdd-btn" @click.stop="onBtnSubmit">Submit</button>
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
                    this.$emit("getRecord")
                    this.close();
                });
        },
        close: function() {
            this.$emit("pageSwitch", false)
        }
    }
};
</script>

<style scoped>
.pageAdd {
    background-color: white;
    cursor: pointer;
    padding: 20px 0;
}

.pageAdd-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin: auto;
    max-width: 500px;
}

.pageAdd-area {
    width: 100%;
}

.pageAdd-btn {
    padding: 10px 20px;
    background-color: blueviolet;
    color: white;
    cursor: pointer;
    margin: 10px 0;
    outline: none;
}

.pageAdd-error {
    padding: 10px 0;
    color: red;
}
</style>

