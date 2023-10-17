package services

import (
	"strings"

	"github.com/cr1m3s/tch_backend/configs"
)

type Email struct {
	message string
}

func NewEmail(recepient string, token string) Email {
	template := Email{message: emailTemplate}

	template.message = strings.Replace(template.message, "{USERNAME}", recepient, 1)
	template.message = strings.Replace(template.message, "{TOKEN}", token, 1)
	template.message = strings.Replace(template.message, "{PAGE}", configs.PASSWORD_RESET_REDIRECT_PAGE, 3)

	return template
}

var emailTemplate string = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Reset password</title>
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Nunito+Sans:opsz,wght@6..12,600&display=swap" rel="stylesheet">    
    <style>
        body {
            width: 320px;
            font-family: Nunito Sans;
            font-weight: 600;
            font-size: 20px;
            line-height: 24px;
            margin: auto;
            padding: 0;
            width: 100%; 
        }

        h1 {
            margin-top: 0;
            margin-bottom: 25px;
            font-size: 24px;
        }

        p {
            margin: 0;
        }         
        
        a {
            text-decoration: none;
            color: currentColor;
        }
        
        .button {
            font-size: 18px;
            color: #FFFFFF;
            line-height: 12px;
            display: block;
            width: 100%;
            cursor: pointer;
            background-color: #0D5D74;
            border:none;
            border-radius: 10px;
            padding-top: 21.5px;
            padding-bottom: 21.5px;
            text-align: center;
            margin-bottom: 50px;
        }    

        .button:hover,
        .button.focus {
            background-color: #063E55;
        }

        .button:active {
            background-color: #061F3E;
        }        

        header {
            background-color: #0D5D74;
            padding-top: 42.5px;
            padding-bottom: 42.5px;
        }

        main {
            background-color: #FFFFFF;
            padding: 100px 0;
            font-size: 24px; 
            line-height: 28px;
        }

        .container {
            width: 288px;
            margin: 0 auto;
        }

        .let-know-link:hover,
        .let-know-link:focus {
            text-decoration: underline;
        }

        footer {
            background-color: #0D5D74;
            padding-top: 50px;
            padding-bottom: 50px;
            color: #FFFFFF;
            text-align: center;
        }

        .footer-nav {
            display: flex; 
            flex-direction: column;
            gap: 15px;
            font-size: 18px;
        }

        .footer-nav-link:hover,
        .footer-nav-link:focus {
            text-decoration: underline;
        }

        .website-info {
            font-size: 16px;
            display: flex;
            flex-direction: column;
            gap: 10px;
        }

        @media only screen and (min-width: 768px) {
            body {
                max-width: 640px;
            }

            main {
                padding-bottom: 150px;
            }

            .container {
                width: 440px;            
            }

            footer {
                padding-top: 40px;
                padding-bottom: 40px;                
            }

            .footer-nav {
                gap: 30px;
            }            

            .website-info {
                font-size: 18px;
                flex-direction: row;
                justify-content: center; 
                gap: 25px;
            }            
        }        
    </style>
</head>
<body>
    <header>
        <div style="display: flex; justify-content: center;">
            <svg width="217" height="72" viewBox="0 0 218 73" fill="none" xmlns="http://www.w3.org/2000/svg">
                <mask id="mask0_524_5118" style="mask-type:luminance" maskUnits="userSpaceOnUse" x="0" y="1" width="71" height="71">
                    <path d="M70.5 1.5H0.5V71.5H70.5V1.5Z" fill="white"/>
                </mask>
                <g mask="url(#mask0_524_5118)">
                    <path d="M57.6785 21.234C58.8817 20.9116 59.9458 20.2023 60.7064 19.2156C61.4668 18.2288 61.8816 17.0193 61.887 15.7735V15.7587C61.887 14.2614 61.2923 12.8255 60.2333 11.7668C59.1747 10.7081 57.7387 10.1133 56.2416 10.1133H56.2271C54.9813 10.1187 53.7721 10.5337 52.7856 11.2942C51.7991 12.0547 51.0897 13.1185 50.7675 14.3218C49.8452 17.6625 47.8523 20.6085 45.095 22.708C42.3373 24.8074 38.9673 25.9443 35.5016 25.9443C32.0358 25.9443 28.6658 24.8074 25.9082 22.708C23.1507 20.6085 21.158 17.6625 20.2357 14.3218C19.9131 13.1186 19.2038 12.0549 18.2171 11.2946C17.2304 10.5343 16.0211 10.1195 14.7755 10.1142H14.7607C13.2634 10.1142 11.8274 10.709 10.7687 11.7677C9.71001 12.8264 9.11523 14.2624 9.11523 15.7596V15.7745C9.12075 17.0201 9.5357 18.2294 10.2962 19.2159C11.0567 20.2025 12.1205 20.9116 13.3237 21.234C16.6645 22.1563 19.6105 24.149 21.7099 26.9066C23.8093 29.6641 24.9462 33.0342 24.9462 36.4999C24.9462 39.9657 23.8093 43.3356 21.7099 46.0933C19.6105 48.851 16.6645 50.8436 13.3237 51.7658C12.1206 52.0884 11.0568 52.7974 10.2964 53.7839C9.53599 54.7704 9.12106 55.9797 9.11555 57.2251V57.2399C9.11555 58.7373 9.71033 60.173 10.769 61.232C11.8278 62.2906 13.2637 62.8853 14.761 62.8853H14.7758C16.0214 62.88 17.2307 62.4648 18.2173 61.7044C19.2038 60.9438 19.9129 59.8801 20.2354 58.6768C21.1577 55.3361 23.1504 52.3901 25.9079 50.2908C28.6654 48.1911 32.0355 47.0543 35.5013 47.0543C38.967 47.0543 42.337 48.1911 45.0947 50.2908C47.852 52.3901 49.8449 55.3361 50.7672 58.6768C51.0894 59.8801 51.7987 60.9441 52.7853 61.7048C53.7718 62.4651 54.9813 62.8803 56.2271 62.8856H56.2419C57.739 62.8856 59.175 62.2906 60.2337 61.232C61.2923 60.1733 61.8873 58.7373 61.8873 57.2402V57.2251C61.882 55.9793 61.4668 54.7701 60.7064 53.7836C59.9458 52.7971 58.8821 52.088 57.6788 51.7655C54.3381 50.8432 51.3921 48.8507 49.2928 46.093C47.1931 43.3353 46.0563 39.9653 46.0563 36.4996C46.0563 33.0338 47.1931 29.6638 49.2928 26.9063C51.3921 24.1487 54.3381 22.156 57.6788 21.2337L57.6785 21.234Z" fill="url(#paint0_linear_524_5118)"/>
                    <path d="M7.88506 4.84191C8.06664 5.02252 8.30001 5.14205 8.55269 5.18388C8.80537 5.22571 9.06479 5.18776 9.29491 5.07531C11.3945 4.06135 13.7569 3.72364 16.0564 4.10873C18.356 4.49381 20.4795 5.58275 22.1342 7.22538C22.3623 7.45522 22.6706 7.58776 22.9943 7.59516C23.318 7.60257 23.632 7.48426 23.8704 7.26506C23.9888 7.15202 24.0835 7.01645 24.1488 6.8663C24.2141 6.71614 24.2487 6.55446 24.2506 6.39074C24.2525 6.22701 24.2217 6.06457 24.1599 5.91294C24.0981 5.76131 24.0066 5.62357 23.8909 5.50778C21.8731 3.49493 19.2788 2.15995 16.468 1.68805C13.6573 1.21614 10.7693 1.63071 8.20476 2.87423C8.03017 2.9608 7.87869 3.08777 7.76297 3.24458C7.64724 3.4014 7.57058 3.58355 7.5393 3.77592C7.50805 3.9683 7.52311 4.16535 7.58324 4.35074C7.64337 4.53613 7.74684 4.70451 7.88506 4.84191Z" fill="url(#paint1_linear_524_5118)"/>
                    <path d="M1.87849 9.19485C0.636489 11.7511 0.218381 14.6298 0.681855 17.4338C1.14533 20.2378 2.46754 22.8288 4.46607 24.8494C4.67623 25.0703 4.95924 25.2076 5.26279 25.2359C5.5663 25.2642 5.86982 25.1816 6.11721 25.0035C6.25734 24.8976 6.37316 24.7629 6.45682 24.6084C6.54048 24.454 6.59006 24.2833 6.60215 24.1081C6.61428 23.9329 6.58867 23.7571 6.52709 23.5926C6.46548 23.4281 6.36932 23.2787 6.24509 23.1545C4.59405 21.4995 3.49858 19.3719 3.11047 17.0666C2.72236 14.7612 3.06082 12.3922 4.07896 10.2878C4.19226 10.0575 4.23084 9.79747 4.18933 9.54414C4.14785 9.29076 4.02831 9.0567 3.84745 8.87451C3.70993 8.73617 3.54143 8.63264 3.35589 8.57248C3.17036 8.51235 2.97316 8.49729 2.78064 8.5286C2.58813 8.55994 2.40587 8.6367 2.24897 8.75259C2.09207 8.86844 1.96506 9.02007 1.87849 9.19485Z" fill="url(#paint2_linear_524_5118)"/>
                    <path d="M61.7077 67.9259C59.6074 68.9402 57.2444 69.2778 54.9441 68.892C52.6441 68.5064 50.5202 67.4166 48.8656 65.773C48.6379 65.5443 48.3307 65.4124 48.0082 65.4051C47.6854 65.3976 47.3726 65.5154 47.1348 65.7337C47.0157 65.8464 46.9206 65.9822 46.8547 66.1324C46.7892 66.2827 46.7543 66.4446 46.7521 66.6087C46.7499 66.7728 46.7807 66.9356 46.8425 67.0874C46.9042 67.2396 46.9959 67.3775 47.1118 67.4935C49.1295 69.5062 51.7237 70.8411 54.5346 71.3132C57.3452 71.7851 60.2333 71.3706 62.7979 70.127C62.9724 70.0404 63.1239 69.9135 63.2398 69.7566C63.3554 69.5997 63.4319 69.4177 63.4631 69.2252C63.4943 69.0328 63.4792 68.8356 63.419 68.6504C63.3589 68.4649 63.2552 68.2967 63.1169 68.1593C62.9355 67.9788 62.7021 67.8595 62.4495 67.8176C62.1972 67.7757 61.9377 67.8135 61.7077 67.9259Z" fill="url(#paint3_linear_524_5118)"/>
                    <path d="M64.7332 48.1333C64.5149 48.3711 64.3971 48.6839 64.4046 49.0067C64.4119 49.3293 64.5439 49.6367 64.7725 49.8641C66.4174 51.52 67.5075 53.6458 67.8921 55.9477C68.277 58.2499 67.9375 60.6144 66.9207 62.7154C66.808 62.9456 66.7698 63.2052 66.8114 63.4581C66.8533 63.711 66.9727 63.9447 67.1535 64.1265C67.2908 64.2648 67.459 64.3681 67.6442 64.4285C67.8294 64.4887 68.0263 64.5038 68.2188 64.4726C68.4109 64.4415 68.5933 64.3652 68.7498 64.2496C68.9067 64.134 69.0339 63.9828 69.1205 63.8083C70.3679 61.2431 70.7852 58.3532 70.3146 55.5401C69.8441 52.7267 68.5089 50.1299 66.4945 48.1103C66.3786 47.994 66.2404 47.9024 66.0882 47.8403C65.9361 47.7783 65.7732 47.7474 65.6091 47.7496C65.4447 47.7518 65.2828 47.7868 65.1322 47.8529C64.9817 47.9188 64.8462 48.0139 64.7332 48.1333Z" fill="url(#paint4_linear_524_5118)"/>
                    <path d="M66.9229 10.2901C67.9387 12.3906 68.2773 14.7547 67.8921 17.0558C67.5069 19.357 66.4167 21.4819 64.7725 23.1373C64.5439 23.3648 64.4119 23.672 64.4046 23.9946C64.3971 24.3172 64.5149 24.6301 64.7332 24.8678C64.8459 24.9868 64.9817 25.082 65.1319 25.1478C65.2822 25.2135 65.4441 25.2485 65.6082 25.2507C65.7723 25.2529 65.9348 25.2222 66.087 25.1604C66.2388 25.0986 66.3771 25.0071 66.493 24.8911C68.5067 22.8723 69.8422 20.2767 70.3134 17.4645C70.7849 14.6523 70.3691 11.7631 69.1237 9.19806C69.0371 9.02353 68.9098 8.87221 68.753 8.75661C68.5961 8.64098 68.414 8.56444 68.2216 8.53332C68.0295 8.5022 67.8323 8.51735 67.6471 8.57758C67.4615 8.6378 67.2933 8.74134 67.156 8.87961C66.9755 9.06133 66.8558 9.29482 66.8139 9.54759C66.7724 9.8004 66.8105 10.0599 66.9229 10.2901Z" fill="url(#paint5_linear_524_5118)"/>
                    <path d="M48.8636 7.22807C50.5189 5.58364 52.6437 4.49358 54.945 4.10835C57.2463 3.72311 59.6102 4.0617 61.7108 5.07737C61.9411 5.1901 62.2006 5.22825 62.4532 5.18651C62.7061 5.14481 62.9396 5.02527 63.1213 4.8446C63.2596 4.7073 63.3632 4.53899 63.4234 4.35367C63.4835 4.16834 63.4986 3.97132 63.4675 3.77896C63.4366 3.58659 63.36 3.40443 63.2445 3.24755C63.1289 3.09068 62.9773 2.96363 62.8028 2.87692C60.2376 1.63146 57.3486 1.21557 54.5365 1.68699C51.7243 2.15843 49.1285 3.49378 47.1098 5.50764C46.9939 5.62365 46.9022 5.76167 46.8405 5.91362C46.7788 6.0656 46.7479 6.22838 46.7501 6.39239C46.7523 6.55643 46.7873 6.71837 46.8528 6.86864C46.9186 7.01895 47.0137 7.15455 47.1328 7.26744C47.3706 7.4856 47.6834 7.60331 48.0063 7.59597C48.3288 7.58863 48.6359 7.45681 48.8636 7.22807Z" fill="url(#paint6_linear_524_5118)"/>
                    <path d="M22.1373 65.7723C20.482 67.4165 18.3571 68.5067 16.0559 68.8919C13.7546 69.2771 11.3906 68.9385 9.29009 67.9227C9.05994 67.8103 8.80039 67.7722 8.54762 67.8137C8.29481 67.8556 8.06132 67.9753 7.8796 68.1558C7.74136 68.2931 7.63782 68.4613 7.5776 68.6469C7.51738 68.8321 7.50219 69.0293 7.53331 69.2214C7.56447 69.4138 7.64101 69.5959 7.7566 69.7528C7.8722 69.9096 8.02355 70.0369 8.19805 70.1235C10.7631 71.3689 13.6523 71.7847 16.4645 71.3132C19.2767 70.842 21.8724 69.5064 23.8911 67.4928C24.0071 67.3769 24.0986 67.2386 24.1604 67.0868C24.2222 66.9349 24.2529 66.7721 24.2507 66.608C24.2485 66.4439 24.2135 66.282 24.1478 66.1317C24.082 65.9815 23.9868 65.8457 23.8678 65.733C23.6301 65.5147 23.3172 65.3969 22.9946 65.4041C22.672 65.4117 22.3647 65.5433 22.1373 65.7723Z" fill="url(#paint7_linear_524_5118)"/>
                    <path d="M4.07771 62.7113C3.06189 60.6107 2.7232 58.2467 3.10838 55.9455C3.49356 53.6442 4.58363 51.5194 6.22809 49.8641C6.4568 49.6364 6.58865 49.3293 6.59599 49.0067C6.60333 48.6839 6.48559 48.3711 6.26746 48.1333C6.15454 48.0142 6.01894 47.9191 5.86866 47.8533C5.71835 47.7878 5.55642 47.7528 5.39241 47.7506C5.2284 47.7484 5.06558 47.7793 4.91364 47.841C4.76169 47.9027 4.62363 47.9944 4.50766 48.1103C2.49379 50.129 1.15842 52.7248 0.686995 55.537C0.215565 58.3491 0.631457 61.2381 1.87692 63.8033C1.96362 63.9778 2.09068 64.1293 2.24755 64.2449C2.40443 64.3605 2.58659 64.4371 2.77896 64.4679C2.97132 64.4991 3.16834 64.484 3.35367 64.4239C3.539 64.3637 3.70729 64.2601 3.84462 64.1218C4.02526 63.94 4.14477 63.7066 4.18656 63.4537C4.22833 63.2011 4.19028 62.9415 4.07771 62.7113Z" fill="url(#paint8_linear_524_5118)"/>
                </g>
                <path d="M112.106 19.844C112.106 21.38 111.522 22.58 110.354 23.444C109.186 24.308 107.354 24.74 104.858 24.74C101.914 24.74 99.7777 24.172 98.4497 23.036C97.1377 21.884 96.4817 20.196 96.4817 17.972H100.322C100.322 19.252 100.666 20.14 101.354 20.636C102.058 21.132 103.266 21.38 104.978 21.38C107.17 21.38 108.266 20.884 108.266 19.892C108.266 19.54 108.074 19.22 107.69 18.932C107.306 18.628 106.898 18.396 106.466 18.236C106.034 18.06 105.362 17.82 104.45 17.516C104.354 17.484 104.274 17.46 104.21 17.444C104.146 17.428 104.066 17.404 103.97 17.372C103.89 17.34 103.81 17.308 103.73 17.276C102.93 17.004 102.346 16.804 101.978 16.676C101.61 16.532 101.114 16.332 100.49 16.076C99.8657 15.804 99.4017 15.556 99.0977 15.332C98.8097 15.092 98.4977 14.796 98.1617 14.444C97.8257 14.092 97.5857 13.7 97.4417 13.268C97.3137 12.836 97.2497 12.348 97.2497 11.804C97.2497 10.364 97.8097 9.26 98.9297 8.492C100.066 7.724 101.722 7.34 103.898 7.34C106.378 7.34 108.25 7.9 109.514 9.02C110.778 10.124 111.41 11.564 111.41 13.34H107.57C107.57 11.58 106.346 10.7 103.898 10.7C102.026 10.7 101.09 11.116 101.09 11.948C101.09 12.332 101.274 12.644 101.642 12.884C102.01 13.124 102.866 13.468 104.21 13.916C106.45 14.684 107.954 15.252 108.722 15.62C110.978 16.756 112.106 18.164 112.106 19.844ZM135.366 7.58L125.166 16.052L135.366 24.5H129.75L122.67 18.524V24.5H118.83V7.58H122.67V13.58L129.75 7.58H135.366ZM153.121 10.94H149.281V21.14H153.121V24.5H141.601V21.14H145.441V10.94H141.601V7.58H153.121V10.94ZM162.158 24.74C161.502 24.756 160.99 24.588 160.622 24.236C160.27 23.884 160.094 23.38 160.094 22.724V7.58H163.934V21.38L172.166 21.14V24.5L162.158 24.74ZM180.95 24.74C180.294 24.756 179.782 24.588 179.414 24.236C179.062 23.884 178.886 23.38 178.886 22.724V7.58H182.726V21.38L190.958 21.14V24.5L180.95 24.74ZM97.9729 55.844C97.9729 57.38 97.3889 58.58 96.2209 59.444C95.0529 60.308 93.2209 60.74 90.7249 60.74C87.7809 60.74 85.6449 60.172 84.3169 59.036C83.0049 57.884 82.3489 56.196 82.3489 53.972H86.1889C86.1889 55.252 86.5329 56.14 87.2209 56.636C87.9249 57.132 89.1329 57.38 90.8449 57.38C93.0369 57.38 94.1329 56.884 94.1329 55.892C94.1329 55.54 93.9409 55.22 93.5569 54.932C93.1729 54.628 92.7649 54.396 92.3329 54.236C91.9009 54.06 91.2289 53.82 90.3169 53.516C90.2209 53.484 90.1409 53.46 90.0769 53.444C90.0129 53.428 89.9329 53.404 89.8369 53.372C89.7569 53.34 89.6769 53.308 89.5969 53.276C88.7969 53.004 88.2129 52.804 87.8449 52.676C87.4769 52.532 86.9809 52.332 86.3569 52.076C85.7329 51.804 85.2689 51.556 84.9649 51.332C84.6769 51.092 84.3649 50.796 84.0289 50.444C83.6929 50.092 83.4529 49.7 83.3089 49.268C83.1809 48.836 83.1169 48.348 83.1169 47.804C83.1169 46.364 83.6769 45.26 84.7969 44.492C85.9329 43.724 87.5889 43.34 89.7649 43.34C92.2449 43.34 94.1169 43.9 95.3809 45.02C96.6449 46.124 97.2769 47.564 97.2769 49.34H93.4369C93.4369 47.58 92.2129 46.7 89.7649 46.7C87.8929 46.7 86.9569 47.116 86.9569 47.948C86.9569 48.332 87.1409 48.644 87.5089 48.884C87.8769 49.124 88.7329 49.468 90.0769 49.916C92.3169 50.684 93.8209 51.252 94.5889 51.62C96.8449 52.756 97.9729 54.164 97.9729 55.844ZM117.428 43.58V46.94H111.908V60.5H108.068V46.94H102.524V43.58H117.428ZM123.911 60.5V43.58H130.463C135.487 43.58 137.999 45.156 137.999 48.308C137.999 50.5 136.567 51.764 133.703 52.1V53.06H136.175C137.535 53.06 138.215 53.74 138.215 55.1V60.5H134.375V54.548H127.751V60.5H123.911ZM127.751 51.188H130.463C131.055 51.188 131.479 51.188 131.735 51.188C132.007 51.172 132.335 51.132 132.719 51.068C133.103 50.988 133.375 50.876 133.535 50.732C133.695 50.588 133.839 50.38 133.967 50.108C134.095 49.836 134.159 49.492 134.159 49.076C134.159 48.66 134.095 48.316 133.967 48.044C133.839 47.756 133.695 47.54 133.535 47.396C133.375 47.252 133.103 47.148 132.719 47.084C132.335 47.004 132.007 46.964 131.735 46.964C131.479 46.948 131.055 46.94 130.463 46.94H127.751V51.188ZM145.422 60.5V43.58H157.494V46.94H149.262V50.348H156.294V53.708H149.262V57.14H157.494V60.5H145.422ZM177.523 60.5L175.987 55.076H169.123L167.611 60.5H163.723L168.595 43.58H176.515L181.411 60.5H177.523ZM170.059 51.716H175.051L173.035 44.54H172.075L170.059 51.716ZM187.905 60.5V43.58H195.465L197.937 59.54H198.897L201.369 43.58H208.929V60.5H205.089V46.94L205.809 44.54H204.849L202.137 60.5H194.697L191.985 44.54H191.025L191.745 46.94V60.5H187.905Z" fill="url(#paint9_linear_524_5118)"/>
                <defs>
                    <linearGradient id="paint0_linear_524_5118" x1="35.5013" y1="67.6878" x2="35.5013" y2="26.5255" gradientUnits="userSpaceOnUse">
                        <stop offset="0.005" stop-color="#3F51DC"/>
                        <stop offset="1" stop-color="#2AC3D1"/>
                    </linearGradient>
                    <linearGradient id="paint1_linear_524_5118" x1="15.887" y1="67.7054" x2="15.887" y2="26.5326" gradientUnits="userSpaceOnUse">
                        <stop offset="0.005" stop-color="#3F51DC"/>
                        <stop offset="1" stop-color="#2AC3D1"/>
                    </linearGradient>
                    <linearGradient id="paint2_linear_524_5118" x1="3.55157" y1="67.6819" x2="3.55157" y2="26.5294" gradientUnits="userSpaceOnUse">
                        <stop offset="0.005" stop-color="#3F51DC"/>
                        <stop offset="1" stop-color="#2AC3D1"/>
                    </linearGradient>
                    <linearGradient id="paint3_linear_524_5118" x1="55.1155" y1="67.6853" x2="55.1155" y2="26.518" gradientUnits="userSpaceOnUse">
                        <stop offset="0.005" stop-color="#3F51DC"/>
                        <stop offset="1" stop-color="#2AC3D1"/>
                    </linearGradient>
                    <linearGradient id="paint4_linear_524_5118" x1="67.4533" y1="67.6855" x2="67.4533" y2="26.5248" gradientUnits="userSpaceOnUse">
                        <stop offset="0.005" stop-color="#3F51DC"/>
                        <stop offset="1" stop-color="#2AC3D1"/>
                    </linearGradient>
                    <linearGradient id="paint5_linear_524_5118" x1="67.4533" y1="67.6864" x2="67.4533" y2="26.5225" gradientUnits="userSpaceOnUse">
                        <stop offset="0.005" stop-color="#3F51DC"/>
                        <stop offset="1" stop-color="#2AC3D1"/>
                    </linearGradient>
                    <linearGradient id="paint6_linear_524_5118" x1="55.1167" y1="67.7068" x2="55.1167" y2="26.5314" gradientUnits="userSpaceOnUse">
                        <stop offset="0.005" stop-color="#3F51DC"/>
                        <stop offset="1" stop-color="#2AC3D1"/>
                    </linearGradient>
                    <linearGradient id="paint7_linear_524_5118" x1="15.8842" y1="67.6846" x2="15.8842" y2="26.5148" gradientUnits="userSpaceOnUse">
                        <stop offset="0.005" stop-color="#3F51DC"/>
                        <stop offset="1" stop-color="#2AC3D1"/>
                    </linearGradient>
                    <linearGradient id="paint8_linear_524_5118" x1="3.54718" y1="67.6798" x2="3.54718" y2="26.5325" gradientUnits="userSpaceOnUse">
                        <stop offset="0.005" stop-color="#3F51DC"/>
                        <stop offset="1" stop-color="#2AC3D1"/>
                    </linearGradient>
                    <linearGradient id="paint9_linear_524_5118" x1="29" y1="-15.9999" x2="135.812" y2="-70.607" gradientUnits="userSpaceOnUse">
                        <stop offset="0.23148" stop-color="#EFCC13"/>
                        <stop offset="1" stop-color="#2AC3D1"/>
                    </linearGradient>
                </defs>
            </svg>
        </div>
    </header>
    <main>
        <div class="container">
            <h1>Hi, <span id="username">{USERNAME}</span>!
            </h1>
            <p style="margin-bottom: 50px;">We got a request to reset your Skill Stream password.</p>
            <a class="button" href="{PAGE}create-password/?token={TOKEN}">Reset password</a>
            <p>
                If you ignore this message, your password will not be changed. If you didn't request a password reset,  <a class="let-know-link" style="color: #338AF3;" href="mailto:exemple@gmail.com">let us know</a>.
            </p>
        </div>

    </main>
    <footer>
        <div style="border-bottom: 1px solid rgba(255, 255, 255, 0.30); margin-bottom: 25px;">
            <p style="margin-bottom: 20px; text-align: center;">Where Education Flows, Knowledge Grows</p>
        </div>
        
        <div class="footer-nav">
            <div style="display: flex; justify-content: center; gap: 25px;">
                <a class="footer-nav-link" href="{PAGE}policy" target="_blank" rel="noreferrer">Privacy Policy</a>
                <a class="footer-nav-link" href="{PAGE}conditions" target="_blank" rel="noreferrer">Terms of Use</a>
            </div>
            <div class="website-info">
                <span>© Skill Stream | All rights reserved</span>
                <span style="display: flex; align-items: center; justify-content: center; gap: 10px">
                    <span>Proudly made in Ukraine</span>
                    <svg width="25" height="16" viewBox="0 0 25 16" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path d="M0.5 0.000244141H24.5V15.9999H0.5V0.000244141Z" fill="#FFDA44"/>
                        <path d="M0.5 0.000244141H24.5V8.00007H0.5V0.000244141Z" fill="#338AF3"/>
                    </svg>
                </span>
            </div>
        </div>
    </footer>
</body>
</html>`
