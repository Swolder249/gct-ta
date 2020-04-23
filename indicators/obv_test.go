package indicators

import (
	"testing"
)

func TestOBV(t *testing.T) {
	expectedOutput := []float64{0, -6710.77, -10905.36, -12409.890000000001, -15526.640000000001, -11502.070000000002, -9922.37, -7338.52, -11061.43, -13700.12, -12581.01, -16553.72, -8480.990000000002, -5224.250000000002, -2516.980000000002, 4211.299999999997, 16369.619999999997, 4456.389999999998, -1500.8500000000022, 7499.139999999998, 3710.8099999999977, 5766.899999999998, 1755.4599999999978, 18764.89, 10158.539999999999, 3445.279999999999, 12808.439999999999, 16486.449999999997, 7700.999999999996, 3280.7999999999965, 6675.489999999996, 3531.679999999996, -5079.330000000004, 1351.0999999999967, -1461.6900000000032, 4878.899999999997, 14498.669999999998, 24064.229999999996, 17582.309999999998, 25527.469999999998, 20517.85, 22186.85, 19171.059999999998, 14412.929999999997, 9545.699999999997, 17609.67, 25164.159999999996, 29177.649999999994, 32446.389999999992, 37276.729999999996, 29716.059999999998, 37843.659999999996, 44397.06999999999, 35049.619999999995, 41239.99999999999, 35477.689999999995, 39611.17, 33697.97, 40452.32, 30515.5, 37123.99, 41395.46, 39298.9, 42628.12, 36751.11, 30471.11, 18906.91, 28180.809999999998, 21633.829999999998, 19295.76, 23452.329999999998, 29639.19, 25009.219999999998, 21655.989999999998, 28104.679999999997, 32241.1, 29435.51, 20366.62, 8189.23, -85.72000000000116, 7203.8899999999985, -51309.5, 3110.1399999999994, -12260.730000000001, 4328.569999999998, -30576.89, -13059.05, 9500.18, 34283.91, 64366.41, 52569.68000000001, 41756.30000000001, 60810.100000000006, 80529.19, 63853.45, 71247.16, 60430.87, 48817.380000000005, 39518.44, 40067.060000000005}

	ret := OBV(testClose, testVolume)
	if len(ret) != len(expectedOutput) {
		t.Fatalf("unexpected length of return slice %v", len(ret))
	}

	for x := range ret {
		if ret[x] != expectedOutput[x] {
			t.Fatalf("unexpected value returned %v", ret[x])
		}
	}
}