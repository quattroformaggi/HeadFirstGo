package prose
import "testing"

func TestTwoElements(t *testing.T){
	list := []string{"яблоко","апельсин"}
	if JoinWithCommas(list) != "яблоко и апельсин"{
		t.Error("не соответствует ожидаемому значению")
	}
}

func TestThreeElements(t *testing.T){
	list := []string{"яблоко","апельсин","груша"}
	if JoinWithCommas(list) != "яблоко, апельсин, и груша"{
		t.Error("не соответствует ожидаемому значению")
	}
}