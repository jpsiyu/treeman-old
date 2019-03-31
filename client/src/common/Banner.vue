<template>
    <div class="banner-frame">
        <div class="banner">
            <div class="grid"></div>
            <div class="grid">
                <img class="logo" src="../assets/images/treeman.png">
            </div>
            <div class="grid">
                <div class="add" @click="activePage(true)">
                    <p class="noselect">+</p>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { MacroEvent } from "../macro";
export default {
    name: "Banner",
    props: ["activePage"]
};
</script>

<style scoped>
.banner-frame {
    width: 100%;
    height: 50px;
    position: relative;
}

.banner {
    height: 50px;
    width: 100%;
    background-color: white;
    box-shadow: 0 2px 1px rgba(0, 0, 0, 0.2);
    margin-bottom: 10px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    position: fixed;
    z-index: 999;
}

.logo {
    width: auto;
    height: 60%;
}

.grid {
    flex: 1;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
}

.grid:nth-child(3) {
    justify-content: flex-end;
}

.add {
    font-size: 20px;
    font-weight: bold;
    width: 30px;
    height: 30px;
    display: flex;
    justify-content: center;
    align-items: center;
    border-radius: 50%;
    background-color: rgb(0, 158, 248);
    color: white;
    margin-right: 20%;
    cursor: pointer;
}
</style>

