<template>
    <div id="app">
        <PagePerson/>
    </div>
</template>

<script>
import PagePerson from "../person/PagePerson";
export default {
    name: "App",
    components: { PagePerson },
    created: function() {}
};
</script>

<style>
body {
    padding: 0;
    margin: 0;
    box-sizing: border-box;
    background-color: rgba(229, 236, 240, 1);
}

.noselect {
    -webkit-touch-callout: none; /* iOS Safari */
    -webkit-user-select: none; /* Safari */
    -khtml-user-select: none; /* Konqueror HTML */
    -moz-user-select: none; /* Firefox */
    -ms-user-select: none; /* Internet Explorer/Edge */
    user-select: none; /* Non-prefixed version, currently*/
}
</style>
