package tcg

var pixelChars2x4Braille = []rune{
	// symbol // code - graphics bits
	'⠀', // 0x2800 - 00000000
	'⢀', // 0x2880 - 00000001
	'⡀', // 0x2840 - 00000010
	'⣀', // 0x28C0 - 00000011
	'⠠', // 0x2820 - 00000100
	'⢠', // 0x28A0 - 00000101
	'⡠', // 0x2860 - 00000110
	'⣠', // 0x28E0 - 00000111
	'⠄', // 0x2804 - 00001000
	'⢄', // 0x2884 - 00001001
	'⡄', // 0x2844 - 00001010
	'⣄', // 0x28C4 - 00001011
	'⠤', // 0x2824 - 00001100
	'⢤', // 0x28A4 - 00001101
	'⡤', // 0x2864 - 00001110
	'⣤', // 0x28E4 - 00001111
	'⠐', // 0x2810 - 00010000
	'⢐', // 0x2890 - 00010001
	'⡐', // 0x2850 - 00010010
	'⣐', // 0x28D0 - 00010011
	'⠰', // 0x2830 - 00010100
	'⢰', // 0x28B0 - 00010101
	'⡰', // 0x2870 - 00010110
	'⣰', // 0x28F0 - 00010111
	'⠔', // 0x2814 - 00011000
	'⢔', // 0x2894 - 00011001
	'⡔', // 0x2854 - 00011010
	'⣔', // 0x28D4 - 00011011
	'⠴', // 0x2834 - 00011100
	'⢴', // 0x28B4 - 00011101
	'⡴', // 0x2874 - 00011110
	'⣴', // 0x28F4 - 00011111
	'⠂', // 0x2802 - 00100000
	'⢂', // 0x2882 - 00100001
	'⡂', // 0x2842 - 00100010
	'⣂', // 0x28C2 - 00100011
	'⠢', // 0x2822 - 00100100
	'⢢', // 0x28A2 - 00100101
	'⡢', // 0x2862 - 00100110
	'⣢', // 0x28E2 - 00100111
	'⠆', // 0x2806 - 00101000
	'⢆', // 0x2886 - 00101001
	'⡆', // 0x2846 - 00101010
	'⣆', // 0x28C6 - 00101011
	'⠦', // 0x2826 - 00101100
	'⢦', // 0x28A6 - 00101101
	'⡦', // 0x2866 - 00101110
	'⣦', // 0x28E6 - 00101111
	'⠒', // 0x2812 - 00110000
	'⢒', // 0x2892 - 00110001
	'⡒', // 0x2852 - 00110010
	'⣒', // 0x28D2 - 00110011
	'⠲', // 0x2832 - 00110100
	'⢲', // 0x28B2 - 00110101
	'⡲', // 0x2872 - 00110110
	'⣲', // 0x28F2 - 00110111
	'⠖', // 0x2816 - 00111000
	'⢖', // 0x2896 - 00111001
	'⡖', // 0x2856 - 00111010
	'⣖', // 0x28D6 - 00111011
	'⠶', // 0x2836 - 00111100
	'⢶', // 0x28B6 - 00111101
	'⡶', // 0x2876 - 00111110
	'⣶', // 0x28F6 - 00111111
	'⠈', // 0x2808 - 01000000
	'⢈', // 0x2888 - 01000001
	'⡈', // 0x2848 - 01000010
	'⣈', // 0x28C8 - 01000011
	'⠨', // 0x2828 - 01000100
	'⢨', // 0x28A8 - 01000101
	'⡨', // 0x2868 - 01000110
	'⣨', // 0x28E8 - 01000111
	'⠌', // 0x280C - 01001000
	'⢌', // 0x288C - 01001001
	'⡌', // 0x284C - 01001010
	'⣌', // 0x28CC - 01001011
	'⠬', // 0x282C - 01001100
	'⢬', // 0x28AC - 01001101
	'⡬', // 0x286C - 01001110
	'⣬', // 0x28EC - 01001111
	'⠘', // 0x2818 - 01010000
	'⢘', // 0x2898 - 01010001
	'⡘', // 0x2858 - 01010010
	'⣘', // 0x28D8 - 01010011
	'⠸', // 0x2838 - 01010100
	'⢸', // 0x28B8 - 01010101
	'⡸', // 0x2878 - 01010110
	'⣸', // 0x28F8 - 01010111
	'⠜', // 0x281C - 01011000
	'⢜', // 0x289C - 01011001
	'⡜', // 0x285C - 01011010
	'⣜', // 0x28DC - 01011011
	'⠼', // 0x283C - 01011100
	'⢼', // 0x28BC - 01011101
	'⡼', // 0x287C - 01011110
	'⣼', // 0x28FC - 01011111
	'⠊', // 0x280A - 01100000
	'⢊', // 0x288A - 01100001
	'⡊', // 0x284A - 01100010
	'⣊', // 0x28CA - 01100011
	'⠪', // 0x282A - 01100100
	'⢪', // 0x28AA - 01100101
	'⡪', // 0x286A - 01100110
	'⣪', // 0x28EA - 01100111
	'⠎', // 0x280E - 01101000
	'⢎', // 0x288E - 01101001
	'⡎', // 0x284E - 01101010
	'⣎', // 0x28CE - 01101011
	'⠮', // 0x282E - 01101100
	'⢮', // 0x28AE - 01101101
	'⡮', // 0x286E - 01101110
	'⣮', // 0x28EE - 01101111
	'⠚', // 0x281A - 01110000
	'⢚', // 0x289A - 01110001
	'⡚', // 0x285A - 01110010
	'⣚', // 0x28DA - 01110011
	'⠺', // 0x283A - 01110100
	'⢺', // 0x28BA - 01110101
	'⡺', // 0x287A - 01110110
	'⣺', // 0x28FA - 01110111
	'⠞', // 0x281E - 01111000
	'⢞', // 0x289E - 01111001
	'⡞', // 0x285E - 01111010
	'⣞', // 0x28DE - 01111011
	'⠾', // 0x283E - 01111100
	'⢾', // 0x28BE - 01111101
	'⡾', // 0x287E - 01111110
	'⣾', // 0x28FE - 01111111
	'⠁', // 0x2801 - 10000000
	'⢁', // 0x2881 - 10000001
	'⡁', // 0x2841 - 10000010
	'⣁', // 0x28C1 - 10000011
	'⠡', // 0x2821 - 10000100
	'⢡', // 0x28A1 - 10000101
	'⡡', // 0x2861 - 10000110
	'⣡', // 0x28E1 - 10000111
	'⠅', // 0x2805 - 10001000
	'⢅', // 0x2885 - 10001001
	'⡅', // 0x2845 - 10001010
	'⣅', // 0x28C5 - 10001011
	'⠥', // 0x2825 - 10001100
	'⢥', // 0x28A5 - 10001101
	'⡥', // 0x2865 - 10001110
	'⣥', // 0x28E5 - 10001111
	'⠑', // 0x2811 - 10010000
	'⢑', // 0x2891 - 10010001
	'⡑', // 0x2851 - 10010010
	'⣑', // 0x28D1 - 10010011
	'⠱', // 0x2831 - 10010100
	'⢱', // 0x28B1 - 10010101
	'⡱', // 0x2871 - 10010110
	'⣱', // 0x28F1 - 10010111
	'⠕', // 0x2815 - 10011000
	'⢕', // 0x2895 - 10011001
	'⡕', // 0x2855 - 10011010
	'⣕', // 0x28D5 - 10011011
	'⠵', // 0x2835 - 10011100
	'⢵', // 0x28B5 - 10011101
	'⡵', // 0x2875 - 10011110
	'⣵', // 0x28F5 - 10011111
	'⠃', // 0x2803 - 10100000
	'⢃', // 0x2883 - 10100001
	'⡃', // 0x2843 - 10100010
	'⣃', // 0x28C3 - 10100011
	'⠣', // 0x2823 - 10100100
	'⢣', // 0x28A3 - 10100101
	'⡣', // 0x2863 - 10100110
	'⣣', // 0x28E3 - 10100111
	'⠇', // 0x2807 - 10101000
	'⢇', // 0x2887 - 10101001
	'⡇', // 0x2847 - 10101010
	'⣇', // 0x28C7 - 10101011
	'⠧', // 0x2827 - 10101100
	'⢧', // 0x28A7 - 10101101
	'⡧', // 0x2867 - 10101110
	'⣧', // 0x28E7 - 10101111
	'⠓', // 0x2813 - 10110000
	'⢓', // 0x2893 - 10110001
	'⡓', // 0x2853 - 10110010
	'⣓', // 0x28D3 - 10110011
	'⠳', // 0x2833 - 10110100
	'⢳', // 0x28B3 - 10110101
	'⡳', // 0x2873 - 10110110
	'⣳', // 0x28F3 - 10110111
	'⠗', // 0x2817 - 10111000
	'⢗', // 0x2897 - 10111001
	'⡗', // 0x2857 - 10111010
	'⣗', // 0x28D7 - 10111011
	'⠷', // 0x2837 - 10111100
	'⢷', // 0x28B7 - 10111101
	'⡷', // 0x2877 - 10111110
	'⣷', // 0x28F7 - 10111111
	'⠉', // 0x2809 - 11000000
	'⢉', // 0x2889 - 11000001
	'⡉', // 0x2849 - 11000010
	'⣉', // 0x28C9 - 11000011
	'⠩', // 0x2829 - 11000100
	'⢩', // 0x28A9 - 11000101
	'⡩', // 0x2869 - 11000110
	'⣩', // 0x28E9 - 11000111
	'⠍', // 0x280D - 11001000
	'⢍', // 0x288D - 11001001
	'⡍', // 0x284D - 11001010
	'⣍', // 0x28CD - 11001011
	'⠭', // 0x282D - 11001100
	'⢭', // 0x28AD - 11001101
	'⡭', // 0x286D - 11001110
	'⣭', // 0x28ED - 11001111
	'⠙', // 0x2819 - 11010000
	'⢙', // 0x2899 - 11010001
	'⡙', // 0x2859 - 11010010
	'⣙', // 0x28D9 - 11010011
	'⠹', // 0x2839 - 11010100
	'⢹', // 0x28B9 - 11010101
	'⡹', // 0x2879 - 11010110
	'⣹', // 0x28F9 - 11010111
	'⠝', // 0x281D - 11011000
	'⢝', // 0x289D - 11011001
	'⡝', // 0x285D - 11011010
	'⣝', // 0x28DD - 11011011
	'⠽', // 0x283D - 11011100
	'⢽', // 0x28BD - 11011101
	'⡽', // 0x287D - 11011110
	'⣽', // 0x28FD - 11011111
	'⠋', // 0x280B - 11100000
	'⢋', // 0x288B - 11100001
	'⡋', // 0x284B - 11100010
	'⣋', // 0x28CB - 11100011
	'⠫', // 0x282B - 11100100
	'⢫', // 0x28AB - 11100101
	'⡫', // 0x286B - 11100110
	'⣫', // 0x28EB - 11100111
	'⠏', // 0x280F - 11101000
	'⢏', // 0x288F - 11101001
	'⡏', // 0x284F - 11101010
	'⣏', // 0x28CF - 11101011
	'⠯', // 0x282F - 11101100
	'⢯', // 0x28AF - 11101101
	'⡯', // 0x286F - 11101110
	'⣯', // 0x28EF - 11101111
	'⠛', // 0x281B - 11110000
	'⢛', // 0x289B - 11110001
	'⡛', // 0x285B - 11110010
	'⣛', // 0x28DB - 11110011
	'⠻', // 0x283B - 11110100
	'⢻', // 0x28BB - 11110101
	'⡻', // 0x287B - 11110110
	'⣻', // 0x28FB - 11110111
	'⠟', // 0x281F - 11111000
	'⢟', // 0x289F - 11111001
	'⡟', // 0x285F - 11111010
	'⣟', // 0x28DF - 11111011
	'⠿', // 0x283F - 11111100
	'⢿', // 0x28BF - 11111101
	'⡿', // 0x287F - 11111110
	'⣿', // 0x28FF - 11111111
}
