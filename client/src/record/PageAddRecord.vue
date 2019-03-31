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
import axios from "axios";
import qs from "qs";
export default {
    name: "PageAddRecord",
    props: ["activePage", "id"],
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
            const ajaxData = {
                detail: this.detail,
                comment: this.comment,
                id: this.id
            };
            axios
                .post("/api/addrecord", qs.stringify(ajaxData))
                .then(response => {
                    window.location.reload()
                });
        },
        close: function() {
            this.activePage(false);
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

