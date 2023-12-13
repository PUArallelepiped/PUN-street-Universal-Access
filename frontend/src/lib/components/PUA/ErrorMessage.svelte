<script lang="ts">
    export let errorMsgVisible: boolean = true;
    export let errorMsg: string = 'Email or Password ERROR';
</script>

<div
    class="flex flex-row justify-center gap-2 p-4"
    id="errorMsg"
    style="visibility: {errorMsgVisible ? 'visible' : 'hidden'}"
>
    <div class="flex items-center">
        <svg
            width="20"
            height="20"
            viewBox="0 0 26 26"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
        >
            <g id="Group 10">
                <circle id="Ellipse 215" cx="13" cy="13" r="13" fill="#DD0808" />
                <path
                    id="Vector"
                    d="M8.58099 5.70711C8.19047 5.31658 7.5573 5.31658 7.16678 5.70711L5.70711 7.16678C5.31658 7.5573 5.31658 8.19047 5.70711 8.58099L6.46752 9.3414L9.4002 12.3235C9.78516 12.715 9.78254 13.3436 9.39432 13.7318L6.46752 16.6586L5.72979 17.3656C5.32801 17.7506 5.31785 18.3896 5.70718 18.7872L7.16686 20.2779C7.55613 20.6755 8.19504 20.6788 8.58847 20.2854L9.3414 19.5325L12.3175 16.5564C12.7104 16.1635 13.3481 16.1662 13.7376 16.5623L16.6586 19.5325L17.3582 20.2625C17.7461 20.6672 18.3909 20.6741 18.7873 20.2777L20.2777 18.7873C20.6741 18.3909 20.6672 17.7461 20.2625 17.3582L19.5325 16.6586L16.5623 13.7376C16.1662 13.3481 16.1635 12.7104 16.5564 12.3175L19.5325 9.3414L20.2854 8.58847C20.6788 8.19504 20.6755 7.55613 20.2779 7.16686L18.7872 5.70719C18.3896 5.31785 17.7506 5.32801 17.3656 5.72979L16.6586 6.46752L13.7318 9.39432C13.3436 9.78254 12.715 9.78516 12.3235 9.4002L9.3414 6.46752L8.58099 5.70711Z"
                    fill="white"
                />
            </g>
        </svg>
    </div>
    <div class="text-center text-xl font-semibold leading-relaxed text-red-600">
        {errorMsg}
    </div>
</div>